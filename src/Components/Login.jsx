import React, { useState } from 'react'

const LoginPage = () => {
    const [input, setInput] = useState({})
    const onInputChange = (e) => setInput({
        ...input,
        [e.currentTarget.name]: e.currentTarget.value
    })

    const onSubmit = (e) => {
        e.preventDefault()
        fetch("/login",{
            method: "POST",
            body: JSON.stringify(input),
            headers:{
                "Accept": "application/json",
                "Content-type": "application/json"
            }
        })
    }

    return (
        <form onSubmit={onSubmit}>
            <div>
                <label>
                    Username:
                </label>
                <input type="text" name="username" onChange={onInputChange}/>
            </div>
            <div>
                <label>
                    Password:
                </label>
                <input type="text" name="password" onChange={onInputChange}/>
            </div>
            <input type="submit"/>
        </form>
    )
}
export default LoginPage;

