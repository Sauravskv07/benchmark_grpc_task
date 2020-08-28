/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
    "github.com/golang-collections/go-datastructures/queue"
    pb "github.com/sauravskv07/benchmark_grpc_task/proto"
)

var (
	port       = flag.Int("port", 10000, "The server port")
    num_components = flag.Int("components",10,"The number of components")
    cap_queues = flag.Int("capacity",1000,"The capacity of each queue")
)

type taskServiceServer struct {
	pb.UnimplementedTaskServiceServer
	queues []*queue.RingBuffer // protect queue from multiple reads
}

// GetFeature returns the feature at the given point.
func(s *taskServiceServer) GetTaskPage(ctx context.Context, in *pb.TaskRequest) (*pb.TaskResponse, error){

    page_size:=in.num_tasks
    comp_id:=in.comp_id
    resp:=new(pb.TaskResponse)
    resp.count=0
    taskArrayPtr := make([]*pb.Task, page_size)
    resp.tasks = taskArrayPtr
    for i:=0; i!=0 && queues[comp_id].Empty()==false; i--{
        task, err:=queues[comp_id].Get()
        if err==nil{
            resp.tasks[i]=&Task
        }else{
            return nil,err
        } 
        count++
    }
    resp.count=count
    return resp,nil
}

func (s *taskServiceServer) AddTask(stream pb.TaskService_AddTaskServer) error{
    
    for {
        task, err:=stream.Recv()
        if err==io.EOF(){
            return stream.SendAndClose(&pb.TaskAddAck{
                message:"Done Adding Task",
            })
        }
        if(err!=nil){
            return err
        }
        
        if queues[task.comp_id].Len()==queues[task.comp_id].Cap() {
            return stream.SendAndClose(&pb.TaskAddAck{
                status:"Queue Overflow! No more tasks can be added at this point. Try Later!",
            })
        }

        err = queues[task.comp_id].Put()
        
        if err!=nil{
            return err
        }
    }
}

func (s *taskServiceServer) GetTaskStream(stream pb.TaskService_GetTaskStreamServer) error {
    for{
        task_req,err:=stream.Recv()
        var flag int8
        if err==io.EOF{
            flag=1
        }
        if err!=io.EOF && err!=nil{
            return err
        }
        comp_id:=task_req.comp_id
        page_size:=task_req.num_tasks;

        if(page_size<=-1){
            
            for queues[comp_id].Empty()==false{
            
                task,err:=queues[comp_id].Get()
            
                if err!=nil{
                    
                    return err
                
                }else{
                    
                    err=stream.send(&task)
                    
                    if err!=nil{
                        return err
                    }
                }
            }
            
            return nil
        
        } else{
            for ;(queues[comp_id].Empty()==false && page_size>0);page_size--{

                task,err:=queues[comp_id].Get()
            
                if err!=nil{
                    
                    return err
                
                }else{
                    
                    err=stream.send(&task)
                    
                    if err!=nil{
                        return err
                    }
                }
            }
            
            return nil              
        }
    }
}

func newServer() *taskServiceServer {
	s := &taskServiceServer{queues: make([]*RingBuffer,num_components)}
    
    for i:=0; i<num_components; i++{
        s.queues[i]= newRingBuffer(cap_queues)
    }
	
    return s
}

func main() {
	flag.Parse()
	
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	
    if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
    grpcServer := grpc.NewServer()
	
    pb.RegisterRouteGuideServer(grpcServer, newServer())
	
    grpcServer.Serve(lis)
}

