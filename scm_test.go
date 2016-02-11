package main

import (
    "testing"
    // "fmt"
)

func TestLisp(t *testing.T) {
    TEST := func(expr string, expected_result string) {
        // fmt.Println("Running ... ", expr)
        got := String(eval(read(expr), &globalenv))
        if got != expected_result {
            t.Error("Expected ", expected_result, " got ", got)
        } else {
            // fmt.Println("OK", expr)
        }
    }

    // the 29 unit tests for lis.py
    TEST("(quote (testing 1 (2.1) -3.14e159))", "(testing 1 (2.1) -3.14e+159)")
    TEST("(+ 2 2)", "4")
    TEST("(+ (* 2 100) (* 1 10))", "210")
    TEST("(if (> 6 5) (+ 1 1) (+ 2 2))", "2")
    TEST("(if (< 6 5) (+ 1 1) (+ 2 2))", "4")
    TEST("(define x 3)", "ok")
    TEST("x", "3")
    TEST("(+ x x)", "6")
    TEST("(begin (define x 1) (set! x (+ x 1)) (+ x 1))", "3")
    TEST("((lambda (x) (+ x x)) 5)", "10")
    TEST("(define twice (lambda (x) (* 2 x)))", "ok")
    TEST("(twice 5)", "10")
    TEST("(define compose (lambda (f g) (lambda (x) (f (g x)))))", "ok")
    TEST("((compose list twice) 5)", "(10)")
    TEST("(define repeat (lambda (f) (compose f f)))", "ok")
    TEST("((repeat twice) 5)", "20")
    TEST("((repeat (repeat twice)) 5)", "80")
    TEST("(define fact (lambda (n) (if (<= n 1) 1 (* n (fact (- n 1))))))", "ok")
    TEST("(fact 3)", "6")
    //TEST("(fact 50)", "30414093201713378043612608166064768844377641568960512000000000000")
    TEST("(fact 12)", "4.790016e+08"); // no bignums; this is as far as we go with 32 bits
    TEST("(define abs (lambda (n) ((if (> n 0) + -) 0 n)))", "ok")
    TEST("(list (abs -3) (abs 0) (abs 3))", "(3 0 3)")
    TEST(`(define combine (lambda (f)
              (lambda (x y)
                 (if (null? x) (quote ())
                 (f (list (car x) (car y))
                 ((combine f) (cdr x) (cdr y)))))))`, "ok")
    TEST("(define zip (combine cons))", "ok")
    TEST("(zip (list 1 2 3 4) (list 5 6 7 8))", "((1 5) (2 6) (3 7) (4 8))")
    TEST(`(define riff-shuffle (lambda (deck) (begin
             (define take (lambda (n seq) (if (<= n 0) (quote ()) (cons (car seq) (take (- n 1) (cdr seq))))))
             (define drop (lambda (n seq) (if (<= n 0) seq (drop (- n 1) (cdr seq)))))
             (define mid (lambda (seq) (/ (length seq) 2)))
             ((combine append) (take (mid deck) deck) (drop (mid deck) deck)))))`, "ok")
    TEST("(riff-shuffle (list 1 2 3 4 5 6 7 8))", "(1 5 2 6 3 7 4 8)")
    TEST("((repeat riff-shuffle) (list 1 2 3 4 5 6 7 8))",  "(1 3 5 7 2 4 6 8)")
    TEST("(riff-shuffle (riff-shuffle (riff-shuffle (list 1 2 3 4 5 6 7 8))))", "(1 2 3 4 5 6 7 8)")
}
