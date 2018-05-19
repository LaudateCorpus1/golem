package golem

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"testing"
	"github.com/bountylabs/golem/dicts"
	"reflect"
)

func TestReadBinary(t *testing.T) {
	b, err := dicts.Asset("data/en.gz")
	if err != nil {
		panic(err)
	}
	_ = b
	gzip.NewReader(bytes.NewBuffer(b))
}
func TestUsage(t *testing.T) {
	l, err := New("english")
	if err != nil {
		fmt.Println(err)
	}
	_ = l
	word, err := l.Lemma("agreement")
	fmt.Println(word, err)
}

func TestLemmatizer_Lemma(t *testing.T) {
	l, _ := New("swedish")
	tests := []struct {
		in      string
		out     string
		wantErr bool
	}{
		{"Avtalet", "avtal", false},
		{"avtalets", "avtal", false},
		{"avtalens", "avtal", false},
		{"Avtaletsadlkj", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := l.Lemma(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lemmatizer.Lemma() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.out {
				t.Errorf("Lemmatizer.Lemma() = %v, want %v", got, tt.out)
			}
		})
	}
}

func TestLemmatizer_Lemmas(t *testing.T) {
	l, _ := New("english")
	tests := []struct {
		in      string
		out     []string
		wantErr bool
	}{
		{"dog", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := l.Lemmas(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lemmatizer.Lemma() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("Lemmatizer.Lemma() = %v, want %v", got, tt.out)
			}
		})
	}
}


func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("swedish")
	}
}

func BenchmarkLookup(b *testing.B) {
	l, _ := New("swedish")
	for i := 0; i < b.N; i++ {
		l.Lemma("Avtalet")
	}
}
