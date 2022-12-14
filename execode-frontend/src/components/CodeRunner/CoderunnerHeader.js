import CoderunnerInput from './CoderunnerInput';
import Select from 'react-select';
import { useState } from 'react';

const CoderunnerHeader = (props) => {

    const [selectedOption, setSelectedOption] = useState('');

    props.sendData(selectedOption)

    const codeRunnerHeadingStyle = {
        display: "flex",
        padding: "20px 20px",
        alignItems: "center",
        justifyContent: "flex-start",
        borderBottom: "1px solid #E5E5E5"
    };

    // langOptions will be use to display dropdown available programming language
    // Addition, we will use React Select library to enhance our code/workflow
    const langOptions = [
        { value: 'c', label: 'C' },
        { value: 'cpp', label: 'C++' },
        { value: 'csharp', label: 'C#' },
        { value: 'go', label: 'Go' },
        { value: 'js', label: 'JavaScript' },
        { value: 'kotlin', label: 'Kotlin' },
        { value: 'lua', label: 'Lua' },
        { value: 'python', label: 'Python3' },
        { value: 'python2', label: 'Python2' },
        { value: 'typescript', label: 'TypeScript' },
    ]

    return (
    <header style={codeRunnerHeadingStyle}>
        <CoderunnerInput />
        <div id='lang-dropdown'>
            <label>Language: </label>
            <Select value={langOptions.value} onChange={setSelectedOption} options={langOptions} defaultValue={selectedOption} />
        </div>
    </header>
  )
}

export default CoderunnerHeader