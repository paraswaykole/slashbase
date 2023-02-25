import React, { useRef, useState } from "react"
import styles from './console.module.scss'

type DBConsolePropType = {
}

const DBConsoleFragment = ({ }: DBConsolePropType) => {

    const [input, setInput] = useState("")
    const [output, setOutput] = useState([
        {
            'text': 'input cmd here',
            'cmd': true
        },
        {
            'text': `output here`,
            'cmd': false
        },
    ])


    const confirmInput = () => {
        const newOutput = [...output]
        newOutput.push({
            text: input,
            cmd: true
        })
        setOutput(newOutput)
        setInput('')
    }

    return <React.Fragment>
        <div className={styles.console}>
            {output.map(block => {
                return <OutputBlock block={block} />
            })}
            <PromptInputWithRef onChange={setInput} confirmInput={confirmInput} />
        </div>
    </React.Fragment>
}

export default DBConsoleFragment


const OutputBlock = ({ block }: any) => {
    return <p className={styles.block + " " + (block.cmd ? styles.cmd : "")}>{block.text}</p>
}


const PromptInputWithRef = (props: any) => {

    const defaultValue = useRef("")
    const inputRef = useRef<HTMLParagraphElement>(null)

    const handleInput = (event: any) => {
        if (props.onChange) {
            props.onChange(event.target.textContent)
        }
    }

    const handleKeyUp = (event: React.KeyboardEvent) => {
        if (props.confirmInput && event.key.toLocaleLowerCase() === 'enter') {
            props.confirmInput()
            if (inputRef.current) {
                inputRef.current.innerText = ""
            }
        }
    }

    return <p
        ref={inputRef}
        className={styles.prompt + " " + styles.cmd}
        contentEditable={true}
        onInput={handleInput}
        onKeyUp={handleKeyUp}
        dangerouslySetInnerHTML={{ __html: defaultValue.current }}
    />

}
