import React from 'react'
import '../styles/welcome.css'


function Welcome() {
  return (
    <div className='welcome'>
      <div className='home'>
        <h1 className='welcome-h1'>Welcome to ExeCode!</h1>
        <h4>an online web application providing integrated​ code learning environment​</h4> <br />
        <a className='welcome-btn' href='https://execode-users.auth.us-east-1.amazoncognito.com/login?client_id=5ujg8c9bbiihttb1gdsijj48t0&response_type=code&scope=email+openid+profile&redirect_uri=http://localhost:3001/home'>
          Login / Sign-Up
        </a>
      </div>
    </div>
    

  )
}

export default Welcome