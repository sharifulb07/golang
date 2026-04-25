package main

import (
	"fmt"
	"sync"
	"time"
)

type Image struct {
	ID       int
	FileName string
	Status   string
}

// ----stage 1: upload-------

func upload(images []Image) <-chan Image {
	out := make(chan Image)

	go func(){

		defer close(out)
	for _, image := range images {

		fmt.Printf("\n Image Processing %s \n", image.FileName)

		time.Sleep(500*time.Microsecond)

		image.Status="uploaded"

		out <-image
	}
	}()
	return out 
}

// stage 2: resize ------

func resize(in <- chan Image) <- chan Image{
	out:=make(chan Image)


	go func(){
		defer close(out)
	for img:=range in{

		fmt.Printf("\n Image Resizing... %s", img.FileName)

		time.Sleep(700*time.Millisecond)
		img.Status="Resized"
		out <-img 

	}

}()
	return out 
}



// stage 3: compressed 

func compress( in <- chan Image ) <- chan Image {

	out:=make(chan Image)

	go func (){

		defer close (out)
	for img:=range in{

		fmt.Printf("\n Image compressing .... %s ", img.FileName)

		time.Sleep(900*time.Millisecond)
		img.Status="compressed"
		out <- img 
	}
}()

	return out
}


// -----stage 4: store 


func store( in <-chan Image, wg *sync.WaitGroup){

	defer wg.Done()


	for img:=range in{

		fmt.Printf("\n image storing ... %s ", img.FileName)

		time.Sleep(400*time.Millisecond)
		img.Status="stored"

		fmt.Printf(
			"Finished Image %d -> %s [%s]",
			img.ID,
			img.FileName,
			img.Status,
		)
	}
}


func main() {

	images:=[]Image{
		{ID: 1, FileName: "photo1.img"},
		{ID: 2, FileName: "photo2.img"},
		{ID: 3, FileName: "photo3.img"},
		{ID: 4, FileName: "photo4.img"},
	}


	var wg sync.WaitGroup


	stage1:=upload(images)
	stage2:=resize(stage1)
	stage3:=compress(stage2)

	wg.Add(1)

	go store(stage3, &wg)

	wg.Wait()

	fmt.Println("\n All Image Processing completed ")



}