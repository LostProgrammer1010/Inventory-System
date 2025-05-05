import React, { useState } from "react";
import "./signup.css"


const SignUp: React.FC = () => {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
  });


  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    console.log(formData)
  }



  return (
    <div id="body">
      <h1>Sign Up</h1>
      <form id="form" onSubmit={handleSubmit}>
        <label
          htmlFor="username" >
          Username:
        </label>
        <input
          type="text"
          id="username"
          name="username"
          onChange={handleChange} />
        <label
          htmlFor="email">
          Email:
        </label>
        <input
          type="email"
          id="email"
          name="email"
          onChange={handleChange} />
        <label
          htmlFor="password">
          Password:
        </label>
        <input
          type="password"
          id="password"
          name="password"
          onChange={handleChange} />
        <button type="submit">Submit</button>
      </form>
    </div >
  );
};

export default SignUp;
