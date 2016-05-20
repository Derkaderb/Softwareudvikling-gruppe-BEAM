open System
open System.IO



type Exp =
    | Plus of Exp * Exp
    | Mult of Exp * Exp
    | Num of int


type Question() = class
    let mutable quest = 1
    let mutable list = [1;2;3;4;5]
    let mutable index = 0

    member this.ran() =
        let ran = System.Random()
        let ranQ = ran.Next(1,3)
        quest <- ranQ


    member this.q2() =
        let rnd = System.Random()
        let l = List.init 5 (fun _ -> rnd.Next (100))
        let n = rnd.Next (0,5)
        index <- n
        list <- l

    member this.getQ
        with get() = quest

    member this.getlist
        with get() = list

    member this.getindex
        with get() = index

end


System.Console.Clear()
let player = new Question()
//    player.rip()
//    if ((player.getQ)=1) then 
//        let expGen = Arb.generate<Exp>
//        let ran2 = System.Random()
//        let example = Gen.sample (ran2.Next(1,10)) (ran2.Next(1,4)) expGen
//        printfn "Example 1:\n%A" example
//    else 
player.q2()
let bob = player.getlist
let rip = player.getindex
let ask = "liste a =" + (sprintf "%A" (bob)) + " Hvad står der på index " + string(rip) + " i listen a ?:"
File.WriteAllText("question.txt", ask)
let ans = string((bob).[rip]) 
File.WriteAllText("svar.txt", ans)
//    let ans = System.Console.ReadLine()
//    if ans = string((player.getbob).[player.getbob2]) 
//    then System.Console.WriteLine "correct!"
//    else System.Console.WriteLine "nonono"

