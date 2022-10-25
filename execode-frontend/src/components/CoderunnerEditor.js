import CodeMirror from '@uiw/react-codemirror'

const exampleCode = "console.log('Hello, World');"

function CoderunnerEditor() {
  return (
    <div>
        <CodeMirror 
            value={exampleCode}
            height="750px"
            width='750px'
        />
    </div>
  )
}

export default CoderunnerEditor