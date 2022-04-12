import React from 'react'
import TextField from'@material-ui/core/TextField'
import { Button, Container } from '@material-ui/core'
import { useForm } from 'react-hook-form'

export default function LoginPage() {
  const {register,handleSubmit} = useForm();
  const onSubmit=(data) => console.log(data)
  return (
    <Container maxWidth="sm">
    <form onSubmit={handleSubmit(onSubmit)}>
        <h1>Login Page</h1>
        <TextField variant='outlined' label="Email" fullWidth autoFocus autoComplete='email'></TextField>
    
    <Button variant='contained' color='primary' type='submit'>Login</Button>
    </form>
    </Container>
  )
}