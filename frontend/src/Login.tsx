import { useState } from "react"
import { useNavigate } from "react-router-dom"



export default function Login(){
    const [email, setEmail] = useState("")
const [password, setPassword] = useState("")
const navigate=useNavigate()
const onSubmit=()=>{
    fetch("http://localhost:8080/login", {
        method: 'POST',
        credentials: 'include',
        body : JSON.stringify({"email":email,
        "password":password})
    }).then(res=>{
        if (res.status===200){
            alert("Successfullly Logged in!!!!!")
            navigate("/")
        }
        else{
           res.text().then(t=>{
               alert(t)
           })
        }
    })
}
    return (
        <div>
            Email
            <br></br>
         <input type={"email"} value={email} onChange={e=>{setEmail(e.target.value)}}></input>
         <br></br>
            Password
            <br></br>
         <input type={"password"} value={password} onChange={e=>{setPassword(e.target.value)}}></input>
         <br></br>
         <button onClick={e=>onSubmit()}>Login</button>
        </div>
    ) 
}