package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "strconv"
)


type Lists struct {
    listA []string
    listB []string
}


func (l Lists) toInteger(li []string) []int {

    nl, str := []int{}, ""

    for _, s := range(li) {
        str += s
        ns, _ := strconv.Atoi(s)
        nl = append(nl, ns)
    }

    return nl
}

func getInputs() ([]Lists, error) {

    file, err := os.Open("./input.txt")
    if err != nil {
        return nil, err
    }

    defer file.Close()

    l := []Lists{}

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), "")
        l = append(l, Lists{listA: line[:5] , listB: line[len(line) - 5:]})
    }

    return l, nil
}

func sort(n int, nums, sorted []int) (int, []int, []int) {

    if n == 0 {
        sorted = append(sorted, nums[0])
        return 0, nums, sorted
    }

    s, i := 99999, 0

    for ii, v := range nums{
        if v < s {
            s, i = v, ii
        }
    }

    sorted = append(sorted, s)
    nums = append(nums[:i], nums[i+1:]...)

    return sort(n - 1, nums, sorted)
}


func sum(a, b []int) int {

    s := 0

    for i, numa := range(a) {
        numb := b[i]
        if numa > numb {
            s += numa - numb
        } else {
            s += numb - numa
        }
    }

    return s
}


func concat(nums []int) int {

    str := ""
    for _, n := range nums {
        str += strconv.Itoa(n)
    }

    num, _ := strconv.Atoi(str)

    return num
}

func sumInputs(l []Lists) int {

    ta, tb := []int{}, []int{}

    for _, ls := range l {
        a, b := ls.toInteger(ls.listA), ls.toInteger(ls.listB)

        na, nb := concat(a), concat(b)

        ta, tb = append(ta, na), append(tb, nb)

    }

    _, _, sa := sort(len(ta) - 1, ta, []int{})
    _, _, sb := sort(len(tb) - 1, tb, []int{})

    s :=sum(sa, sb)

    return s
}

func main () {
    inputs, err := getInputs()
    if err != nil {
        return
    }

    sum := sumInputs(inputs)
    output := fmt.Sprintf("[AoC] [Day 01] Total distance between lists : %d", sum)
    fmt.Println(output)
}
