import CodeMirror from '@uiw/react-codemirror'
import CoderunnerBtn from './CoderunnerBtn'

const exampleCode = "console.log('Hello, World');"

function CoderunnerEditor() {
  return (
    <div>
        <CodeMirror 
            value={exampleCode}
            height="750px"
            width='750px'
        />
        <CoderunnerBtn name="Run" />
    </div>
  )
}

export default CoderunnerEditor