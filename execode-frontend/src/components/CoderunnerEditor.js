import CodeMirror from '@uiw/react-codemirror'
import CoderunnerRunBtn from './CoderunnerRunBtn'

const exampleCode = "console.log('Hello, World');"

function CoderunnerEditor() {
  return (
    <div>
        <CodeMirror 
            value={exampleCode}
            height="750px"
            width='750px'
        />
        <CoderunnerRunBtn />
    </div>
  )
}

export default CoderunnerEditor