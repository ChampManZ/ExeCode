import { useRef, useState, useEffect } from "react"

function CoderunnerInput() {
    // State Management
    const [title, setTitle] = useState('Untitled Problem Statement')  // Initial State for Title a.k.a Default Value is Untitled Problem Statement
    const [textWidth, setTextWidth] = useState(0)  // Initial State for width. This keep work reactive for watching text length
    const span = useRef();

    useEffect(() => {
        setTextWidth(span.current.offsetWidth)
    }, [title])

    const inputStyle = {
        padding: "5px 5px",
        margin: "8px 0",
        border: "0",
        fontSize: "16px",
        width: textWidth,
        minWidth: "1px"
    }

    const hideTextStyle = {
        position: "absolute", 
        opacity: "0",
        zIndex: "-100",
        whiteSpace: "pre"
    }

    return (
    <wrapper is="custom">
        <span id="hide" ref={span} style={hideTextStyle}>{title}</span>
        <input 
            type="text" 
            name="probTitle" 
            id="probTitle" 
            style={inputStyle} 
            onChange={event => setTitle(event.target.value)}
            autoFocus
            value={title}
        />
    </wrapper>
    )
}

export default CoderunnerInput