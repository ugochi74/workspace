package main

func Gcd(a,b unit) unit {
     if b == 0 {
        return a
     }
     return Gcd(b,a%b)
}
