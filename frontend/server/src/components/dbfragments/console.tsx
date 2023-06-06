import React, { useContext, useEffect, useRef, useState } from "react"
import { Tab } from "../../data/models"
import { initConsole, runConsoleCmd, selectBlocks } from "../../redux/consoleSlice"
import { selectDBConnection } from "../../redux/dbConnectionSlice"
import { useAppDispatch, useAppSelector } from "../../redux/hooks"
import TabContext from "../layouts/tabcontext"
import styles from './console.module.scss'

type DBConsolePropType = {
}

const DBConsoleFragment = ({ }: DBConsolePropType) => {

    const dispatch = useAppDispatch()

    const currentTab: Tab = useContext(TabContext)!

    const consoleEndRef = useRef<HTMLSpanElement>(null)

    const dbConnection = useAppSelector(selectDBConnection)
    const output = useAppSelector(selectBlocks)
    const [input, setInput] = useState("")
    const [nfocus, setFocus] = useState<number>(0)
    const commands = output.filter( e => e.cmd ===  true)
    const [pointer, setPointer] = useState<number>(commands.length-1)

    useEffect(() => {
        dispatch(initConsole(dbConnection!.id))
    }, [dbConnection])

    useEffect(() => {
        consoleEndRef.current?.scrollIntoView({ behavior: 'smooth' })
    }, [output])

    const confirmInput = () => {
        dispatch(runConsoleCmd({ dbConnId: dbConnection!.id, cmdString: input }))
        setInput('')
    }

    const focus = (e: any) => {
        if (e.target.id === "console") {
            setFocus(Math.random())
        }
    }

    return <div className={styles.console + " " + (currentTab.isActive ? "db-tab-active" : "db-tab")} id="console" onClick={focus}>
        <OutputBlock block={{
            text: "Start typing command and press enter to run it.\nType 'help' for more info on console.",
            cmd: false
        }} />
        {output.map((block, idx) => {
            return <OutputBlock block={block} key={idx} />
        })}
        <PromptInputWithRef onChange={setInput} isActive={currentTab.isActive} nfocus={nfocus} confirmInput={confirmInput} commands={commands} pointer={pointer} setPointer={setPointer} />
        <span ref={consoleEndRef}></span>
    </div>
}

export default DBConsoleFragment


const OutputBlock = ({ block }: any) => {
    return <p className={styles.block + " " + (block.cmd ? styles.cmd : "")}>{block.text}</p>
}


const PromptInputWithRef = (props: any) => {

    const defaultValue = useRef("")
    const inputRef = useRef<HTMLParagraphElement>(null)

    useEffect(() => {
        if (props.isActive) {
            inputRef.current?.focus()
        }
    }, [props.isActive, props.nfocus])

    const handleInput = (event: any) => {
        if (props.onChange) {
            props.onChange(event.target.textContent)
        }
    }

    const setInputRef = ( cmd : string) => {
        if(inputRef.current !== null){
            inputRef.current.textContent = cmd;
        }
    }
    const handleKeyUp = (event: React.KeyboardEvent) => {
        if (props.confirmInput && event.key.toLocaleLowerCase() === 'enter') {
            props.confirmInput()
            if (inputRef.current) {
                inputRef.current.innerText = ""
            }
        }
        if ( event.key.toLocaleLowerCase() === 'arrowup') {
            props.setPointer( () => ((props.pointer + props.commands.length -1 ) % props.commands.length))
            setInputRef(props.commands.at(props.pointer)?.text)
        }
        if ( event.key.toLocaleLowerCase() === 'arrowdown'){
            props.setPointer( () => ((props.pointer + 1 ) % props.commands.length))
            setInputRef(props.commands.at(props.pointer)?.text)
        }
    }

    return <p
        ref={inputRef}
        className={styles.prompt}
        contentEditable={true}
        onInput={handleInput}
        onKeyUp={handleKeyUp}
        spellCheck="false"
        dangerouslySetInnerHTML={{ __html: defaultValue.current }}
    />

}
