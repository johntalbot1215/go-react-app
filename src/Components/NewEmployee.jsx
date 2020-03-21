import React, { useState } from 'react';
export const NewEmployee = () => {
    const [input, setInput] = useState({})
  
    const handleInputChange = (e) => setInput({
      ...input,
      [e.currentTarget.name]: e.currentTarget.value
    })

    const handleSubmit = (e) => {
        e.preventDefault()
        fetch('/new-account', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body: JSON.stringify( input ),
        })
    }   
  
    return (
      <form onSubmit={handleSubmit}>
        <div>
          <label>Username:</label>
          <input type="text" name="username" onChange={handleInputChange} />
        </div>
        <div>
          <label>Password:</label>
          <input type="text" name="password" onChange={handleInputChange} />
        </div>
        <input type="submit" />
      </form>
    )
  }

export default NewEmployee;