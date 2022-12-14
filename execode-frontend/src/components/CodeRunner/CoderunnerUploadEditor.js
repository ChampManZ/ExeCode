import React, { useState } from 'react'
// import axios from 'axios'
import CodeMirror from '@uiw/react-codemirror'
import CoderunnerBtn from './CoderunnerBtn'

const uploadBtnStyle = {
    backgroundColor: "#F9A084",
    color: "white",
    padding: "0.5rem",
    borderRadius: "0.3rem",
    cursor: "pointer",
    marginTop: "1rem"
}

const divPaddingStyle = {
    padding: "20px 2px"
}

function CoderunnerUploadEditor(props) {

    const [isClick, setIsClick] = useState(false)
    const [stateTestCases, setStateTestCases] = useState([])
    const [myLength, setMyLength] = useState(0)
    let testCases = []
    const [code, setCode] = useState("// Write your test case function here")

    const onChange = React.useCallback((value, viewUpdate) => {
      setCode(value)
    }, []);
    
    // const testCaseURL = ""
    const submitTestCase = () => {
        testCases = code.split(";")

        setStateTestCases(testCases)
        setIsClick(false)
        setMyLength(testCases.length)
        console.log(testCases)
        console.log(testCases[0])
    }

    const allTestCase = stateTestCases.map((element, key) => {
        return (
            <div key={key}>
                <p>{element}</p>
            </div>
        )
    })

    return (
        <div>
            { isClick ? 
            <div style={divPaddingStyle}>
                <CodeMirror 
                height="auto"
                width='auto'
                value={code}
                onChange={onChange}
                />
                <CoderunnerBtn name="Add" style={uploadBtnStyle} clickFunc={submitTestCase} />
            </div> :
            <CoderunnerBtn name="Add Test Case" style={uploadBtnStyle} clickFunc={(event) => { setIsClick(true) }} />
            }
            {`Input n = ${myLength}`}
            {allTestCase}
        </div>
    )
}

export default CoderunnerUploadEditor