package fib

import (
	"log"
	"testing"
)

func TestRetrieve10thValueFromFibonacciSequence(t *testing.T) {
	want := "55"
	got, err := RetrieveNthFibonacci(uint16(10))

	if (err != nil) || (want != got) {
		t.Errorf("Got %v (Err == %v); wantted %v", got, err, want)
	}

	log.Printf("\nGot: %v\n", got)
}

func TestRetrieve1474thValueFromFibonacciSequence(t *testing.T) {
	want := "49922546054777481906385350680403592275640935929543766946665460861799397458954450591595810323535012493220473350993776311634367136001785103997860984596538045836852574604944378601304041138240342664654575229340490548643476863768120970913864649785803593822617249920299338778001474632358941465522588251623218216960"
	got, err := RetrieveNthFibonacci(uint16(1474))

	if (err != nil) || (want != got) {
		t.Errorf("Got %v (Err == %v); wantted %v", got, err, want)
	}

	log.Printf("\nGot: %v\n", got)
}

func TestErrorRetrieve1475thValueFromFibonacciSequence(t *testing.T) {
	want := "ERROR: invalid Fibonacci Sequence index position [1475]"
	got, err := RetrieveNthFibonacci(uint16(1475))

	if err == nil {
		t.Errorf("Got %v (Err == %v); wantted an error [%v], but got err == nil", got, err, want)
	}
}

func TestRetrieve72thValueFromFibonacciSequence(t *testing.T) {
	want := "498454011879264"
	got, err := RetrieveNthFibonacci(uint16(72))

	if (err != nil) || (want != got) {
		t.Errorf("Got %v (Err == %v); wantted %v", got, err, want)
	}

	log.Printf("\nGot: %v\n", got)
}

func TestRetrieve0thValueFromFibonacciSequence(t *testing.T) {
	want := "1"
	got, err := RetrieveNthFibonacci(uint16(0))

	if (err != nil) || (want != got) {
		t.Errorf("Got %v (Err == %v); wantted %v", got, err, want)
	}

	log.Printf("\nGot: %v\n", got)
}

func TestRetrieve1thValueFromFibonacciSequence(t *testing.T) {
	want := "1"
	got, err := RetrieveNthFibonacci(uint16(1))

	if (err != nil) || (want != got) {
		t.Errorf("Got %v (Err == %v); wantted %v", got, err, want)
	}

	log.Printf("\nGot: %v\n", got)
}
