package main
import "fmt"


func main() {
    
    var a, b, flag, mdc int
    var contador = 1
    
    fmt.Scan(&a)
    fmt.Scan(&b)
    
    for flag != -1{
        
        if(a % contador == 0 && b % contador == 0){
            
            if(contador > mdc){
                
                mdc = contador
            }
        }
        
        // contador
        contador = contador + 1
        
        if(contador > a || contador > b){
            
            flag = -1
        }
    }
    
    fmt.Println(mdc)
}