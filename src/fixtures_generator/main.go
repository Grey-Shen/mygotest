package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Opts struct {
	bufSize       int
	writerBufSize int
}

var (
	imageOpts = Opts{
		bufSize:       50 * 1024,
		writerBufSize: 256 * 1024,
	}
	videoOpts = Opts{
		bufSize:       2 * 1024 * 1024,
		writerBufSize: 4096,
	}
)

var (
	Number          int
	Type            string
	Output          string
	ImageSamplesDir string
	VideoSamplesDir string
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func init() {
	flag.IntVar(&Number, "n", 30, "Number of files to generate.")
	flag.StringVar(&Type, "t", "image", "Type of the file: \"image\" or \"video\".")
	flag.StringVar(&Output, "o", "output", "Output directory.")
	flag.StringVar(&ImageSamplesDir, "is", "samples/images", "Sample images directory.")
	flag.StringVar(&VideoSamplesDir, "vs", "samples/videos", "Sample videos directory.")
}

func main() {
	flag.Parse()

	os.RemoveAll(Output)
	os.MkdirAll(Output, 0755)

	switch Type {
	case "image":
		{
			generateRandomFiles(ImageSamplesDir, imageOpts)
		}
	case "video":
		{
			generateRandomFiles(VideoSamplesDir, videoOpts)
		}
	default:
		{
			log.Fatalf("Invalid file type: %s", Type)
		}
	}
}

func generateRandomFiles(samplesDir string, opts Opts) {
	samples, err := filepath.Glob(filepath.Join(samplesDir, "*"))
	check(err)
	start := time.Now()
	for i := 0; i < Number/len(samples); i++ {
		sample := samples[i%len(samples)]
		outfile := filepath.Join(Output, fmt.Sprintf("%04d%s", i, filepath.Ext(sample)))
		generateRandomFile(sample, outfile, opts)
	}
	end := time.Since(start)
	fmt.Println("total time:", end)
}

func generateRandomFile(sample string, outfile string, opts Opts) {
	inFile, err := os.Open(sample)
	check(err)
	defer inFile.Close()

	outFile, err := os.Create(outfile)
	check(err)
	defer outFile.Close()

	in := bufio.NewReaderSize(inFile, opts.writerBufSize)
	out := bufio.NewWriterSize(outFile, opts.writerBufSize)
	defer out.Flush()

	buf := make([]byte, opts.bufSize)

	for {
		n, err := in.Read(buf)
		if err == io.EOF {
			break
		}
		check(err)

		rand.Read(buf[opts.bufSize-1024:])

		_, err = out.Write(buf[:n])
		check(err)
	}
}
