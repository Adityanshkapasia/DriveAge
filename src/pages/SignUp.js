import React,{useState} from 'react'

const SignUp = () => {

    const[name, setName]=useState("")
    const[email, setEmail]=useState("")
    const[password, setPassword]=useState("")
    
    function signup()
    {
      
    }
  
    return (

    
    <form>
      <div>
          <label htmlFor='name'> Name</label>
          <input id='name' value={name} onClick={(e)=>setName(e.target.value)} type="text" placeholder='Name'/>
      </div>
      <div>
          <label htmlFor='email'>Email</label>
          <input id='email' value={email} onClick={(e)=>setEmail(e.target.value)} type="text"/>
      </div>
      <div>
          <label htmlFor='email'>Password</label>
          <input id='email' value={password} onClick={(e)=>setPassword(e.target.value)} type="password"/>
      </div>
      <button onClick={signup} type='submit'>Login</button>
    </form>
  )
}

export default SignUp