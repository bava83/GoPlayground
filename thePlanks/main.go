package main

import ("fmt")

func main(){
  max := 4
  head := node{length:0, plank: 0}
  insertChild(&head,max)
//  printNodes(&head)
  result := burtForceSearch(&head)
  fmt.Printf("There are %d ways to get to Lenght of 4\n", result)
}

type node struct{
  length int
  plank int
  childrend []node
}

var planks = []int{1,2,3} 

func insertChild(n *node, max int){
  if n.length==max{
    return
  }
  for _,c := range planks{
    if max - n.length - c >= 0{
      n.childrend = append(n.childrend, node{plank: c, length: n.length+c})
    }   
  }
  for i := range n.childrend{
    insertChild(&n.childrend[i],max)
  }
}

func printNodes(n *node){
  fmt.Printf("current: %d", n.length)  
  fmt.Printf("has child %d \n", len(n.childrend))
  for _, c := range n.childrend{
    printNodes(&c)
  }
}

func burtForceSearch(n *node) int{
  if len(n.childrend) == 0{
    return 1
  }
  result := 0
  for i := range n.childrend{
    result += burtForceSearch(&n.childrend[i])
  }
  return result
}

