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
  textAlign: "right",
  padding: "10px 4px"
}

function CoderunnerEditor(props) {

  const [code, setCode] = useState("console.log('Hello, World')\n")

  const onChange = React.useCallback((value, viewUpdate) => {
    setCode(value)
  }, []);

  const version = {
    c: "10.2.0",
    cpp: "10.2.0",
    csharp: "6.12.0",
    go: "1.16.2",
    js: "16.3.0",
    kotlin: "1.4.31",
    lua: "5.4.2",
    python: "3.10.0",
    python2: "2.7.18",
    typescript: "4.2.3"
  }

  // Don't forget to provide here
  const submissionURL = "https://emkc.org/api/v2/piston/execute"
  const submitCode = () => {
    axios.post(submissionURL, {
      "language": props.data.value,
      "version": version[props.data.value],
      "files": [
          {
              "name": "",
              "content": code
          }
      ],
      "stdin": "",
      "args": props.testCases
    })
    .then((res) => props.outputData(res))
  }

  return (
    <div>
        <CodeMirror className='editor-main'
            value={code}
            height="750px"
            width='auto'
            onChange={onChange}
        />
        <div style={divBtnStyle}>
          <CoderunnerBtn name="Run" style={uploadBtnStyle} clickFunc={submitCode} />
        </div>
    </div>
  )
}

export default CoderunnerEditor