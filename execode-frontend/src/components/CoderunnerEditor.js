import React from 'react'
import axios from 'axios'
import CodeMirror from '@uiw/react-codemirror'
import { useState } from 'react'
import CoderunnerBtn from './CoderunnerBtn'

const uploadBtnStyle = {
  backgroundColor: "green",
  color: "white",
  padding: "0.5rem",
  borderRadius: "0.3rem",
  cursor: "pointer",
  marginTop: "1rem",
}

const divBtnStyle = {
  textAlign: "right"
}

function CoderunnerEditor() {

  const [code, setCode] = useState("console.log('Hello, World')")
  const onChange = React.useCallback((value, viewUpdate) => {
    setCode(value)
  }, []);

  // Don't forget to provide here
  const submissionURL = ""
  const submitCode = () => {
    axios.post(submissionURL, {code}).then((res) => console.log(res))
    console.log(code)
  }

  return (
    <div>
        <CodeMirror 
            value={code}
            height="750px"
            width='750px'
            onChange={onChange}
        />
        <div style={divBtnStyle}>
          <CoderunnerBtn name="Run" style={uploadBtnStyle} clickFunc={submitCode} />
        </div>
    </div>
  )
}

export default CoderunnerEditor