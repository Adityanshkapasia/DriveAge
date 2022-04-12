import React from 'react'
import TextField from'@material-ui/core/TextField'
import { Button, Container } from '@material-ui/core'
import { useForm } from 'react-hook-form'

export default function LoginPage() {
  const {register,
    handleSubmit,
     formState: {errors}
    } = useForm();
  const onSubmit=(data) => console.log(data);
  return (
    <Container maxWidth="sm">
    <form onSubmit={handleSubmit(onSubmit)}>
        <h1>Login Page</h1>
        <TextField 
        variant='outlined' label="Email" fullWidth 
        autoFocus autoComplete='email'
        {...register("email",{required: "Required field",
      pattern: { value:/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
      message: "Invalid email address"}})}
        error={!!errors?.email}
        helperText={errors?.email ? errors.email.message : null }
        >

        </TextField>
    
    <Button variant='contained' color='primary' type='submit'>Login</Button>
    </form>
    </Container>
  )
}