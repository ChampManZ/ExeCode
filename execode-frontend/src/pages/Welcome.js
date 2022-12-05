import React from 'react'

const welcomeStyle = {
    textAlign: "center"
}

function Welcome() {
  return (
    <div style={welcomeStyle}>
        <h1>Welcome to ExeCode!</h1>
        <a href='https://execode-users.auth.us-east-1.amazoncognito.com/login?client_id=5ujg8c9bbiihttb1gdsijj48t0&response_type=code&scope=email+openid+profile&redirect_uri=http://localhost:3001/home'>Press to Login/Sign-Up</a>
    </div>
  )
}

export default Welcome