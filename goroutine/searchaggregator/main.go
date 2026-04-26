package main 

import (
	"math/rand"
	"fmt"
	"context"
	"time"
)

// SearchResult Struct

type SearchResult struct{
	Source string 
	Data string 
	Err error
}


// google

func googleSearch(ctx context.Context, query string )<- chan SearchResult{
	out:=make( chan SearchResult, 1)

	go func(){
		defer close(out)

		time.Sleep(time.Duration(rand.Intn(3)+1)*time.Second)

		select{
		case <-ctx.Done():
			return
		default:
			out <- SearchResult{
				Source: "Google",
				Data: "Google Result for : "+query,
			}
		}


	}()
	return out 
}
// bing

func bingSearch(ctx context.Context, query string) <- chan SearchResult{

	out:=make(chan SearchResult, 1)

	
	
	go func(){
		defer close(out)
		time.Sleep(time.Duration(rand.Intn(3)+1)*time.Second)

		select{
		case <- ctx.Done():
			return

		default:
			out <-SearchResult{
				Source: "Bing",
				Data: "Bing search for ..."+query,
			}
		}


	}()


	return out 
}
// elastic

func elasticSearch(ctx context.Context, query string) <- chan SearchResult{
	out:=make(chan SearchResult, 1)

	
	go func(){
		
		defer close(out)
		time.Sleep(time.Duration(rand.Intn(3)+1)*time.Second)

		select{
		case <-ctx.Done():
			return
		default:
			out<-SearchResult{
				Source: "Elastic Search",
				Data: "Elastic Searching for ..."+query,
			}
		}
	}()

	return out 
}
// database


func databaseSearch(ctx context.Context, query string) <- chan SearchResult{

	out:=make(chan SearchResult, 1)

	
	go func(){
		
		defer close(out)
		time.Sleep(time.Duration(rand.Intn(3)+1)*time.Second)

		select{
		case <- ctx.Done():
			return
		default:
			out <-SearchResult{
				Source: "Database",
				Data: "Database Search for ..."+query,
			}
		}
	}()

	return out 
}
// fanin


func fanin(
	ctx context.Context,
	channels ...<-chan SearchResult,

)<- chan SearchResult{
	out:=make(chan SearchResult)

	for _, ch:= range channels{
		go func(c<- chan SearchResult){

			select{
			case result:= <-c:
				if result.Data!=""{
					out<-result
				}

			case <-ctx.Done():
				return
			}
		}(ch)
	}
	return out 

}
// main


func main(){
	rand.Seed(time.Now().UnixNano())

	query:="golang concurrency"

	ctx, cancel:=context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

	google:=googleSearch(ctx, query)
	bing:=bingSearch(ctx, query)
	elastic:=elasticSearch(ctx, query)
	db:=databaseSearch(ctx, query)


	resultChan:=fanin(ctx, google, bing, elastic, db)

	select{
	case result:=<- resultChan:
		fmt.Println("Fastest Result: \n")
		fmt.Printf("Source: %s \n", result.Source)
		fmt.Printf("Data: %s \n", result.Data)


	case <- ctx.Done():
		fmt.Println("Search Timeout: no fast response ")
	}
}