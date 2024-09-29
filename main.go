package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-30

A golang microbiome and metagenomics data analyzer.
- You can extract operon,
- You can extract defined uORFs,
- You can extract and tokensize the generate tags for the machine learning.
- You can make a neural model for the deep learning.
- You can pass all the metagenomics annotations and generate the genomes and sub-genomes files.

*/

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var (
	gbfile string
	gtf    string
	gff    string
)

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "This is a application for the analysis of the bacterial genomes",
}

var sequenceCmd = &cobra.Command{
	Use:  "gbfile",
	Long: "extracting the gbfile for the sequence information",
	Run:  gbFunc,
}

var gtfCmd = &cobra.Command{
	Use:  "gtffile",
	Long: "extracting the information from the GTF file",
	Run:  gtfFunc,
}

var gffCmd = &cobra.Command{
	Use:  "gfffile",
	Long: "extracting the information from the GFF file",
	Run:  gffFunc,
}

func init() {
	sequenceCmd.Flags().
		StringVarP(&gbfile, "gbfile", "g", "taking the gbfile", "sequence genome addition")
	gtfCmd.Flags().
		StringVarP(&gtf, "gtffile", "t", "preparing the gtf file", "gtf sequence mapping")
	gffCmd.Flags().
		StringVarP(&gff, "gfffile", "f", "annotating the gff file", "annotation of the gff file")

	rootCmd.AddCommand(sequenceCmd)
	rootCmd.AddCommand(gtfCmd)
	rootCmd.AddCommand(gffCmd)
}

func gbFunc(cmd *cobra.Command, args []string) {
	fOpen, err := os.Open(gbfile)
	if err != nil {
		log.Fatal(err)
	}
	fRead := bufio.NewScanner(fOpen)
	seqs := []string{}
	genomeSeq := []string{}
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), "A") || strings.HasPrefix(string(line), "T") ||
			strings.HasPrefix(string(line), "G") ||
			strings.HasPrefix(string(line), "C") {
			seqs = append(seqs, string(line))
		}
		// have to check this line as i am passing a strings.Join on a slice
		genomeSeq = strings.Join(seqs, "")
	}
}

type gtfStruct struct {
	id    string
	start string
	end   string
}

type gffStruct struct {
	id    string
	start string
	end   string
}

// Have to code this and the rest of the functions tomorrow which involves the deep learning ones in Golang.

func gtffile(cmd *cobra.Command, args []string) {
	fOpen, err := os.Open(gtf)
	if err != nil {
		log.Fatal(err)
	}
	fRead := bufio.NewScanner(fOpen)
	storegtF := []gtfStruct{}
	for fRead.Scan() {
	}
}
