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

    const consoleRef = useRef<HTMLDivElement>(null)
    const consoleEndRef = useRef<HTMLDivElement>(null)

    const dbConnection = useAppSelector(selectDBConnection)
    const output = useAppSelector(selectBlocks)
    const [input, setInput] = useState("")
    const [nfocus, setFocus] = useState<number>(0)
    const history = output.filter(e => e.cmd).filter(e => e.text !== "").map(e => e.text)
    useEffect(() => {
        dispatch(initConsole(dbConnection!.id))
    }, [dbConnection])

    useEffect(() => {
        scrollToBottom("smooth")
    }, [output])

    const confirmInput = () => {
        dispatch(runConsoleCmd({ dbConnId: dbConnection!.id, cmdString: input }))
        setInput('')
    }

    const focus = (e: any) => {
        if (consoleRef.current?.contains(e.target)) {
            setFocus(Math.random())
        }
    }

    const scrollToBottom = (behavior: ScrollBehavior) => {
        const mainContentDiv = consoleRef.current?.parentNode as HTMLDivElement
        if (mainContentDiv.scrollTop !== consoleEndRef.current?.offsetTop)
            mainContentDiv.scrollTo({ top: consoleEndRef.current?.offsetTop, behavior })
    }

    return <div className={styles.console + " " + (currentTab.isActive ? "db-tab-active" : "db-tab")} id="console" ref={consoleRef} onClick={focus}>
        <OutputBlock block={{
            text: "Start typing command and press enter to run it.\nType 'help' for more info on console.",
            cmd: false
        }} />
        {output.map((block, idx) => {
            return <OutputBlock block={block} key={idx} />
        })}
        <PromptInputWithRef onChange={setInput} isActive={currentTab.isActive} nfocus={nfocus} scrollToBottom={scrollToBottom} confirmInput={confirmInput} history={history} />
        <div id="consoleend" className={styles.consoleend} ref={consoleEndRef}></div>
    </div>
}

export default DBConsoleFragment


const OutputBlock = ({ block }: any) => {
    return <p className={styles.block + " " + (block.cmd ? styles.cmd : "")}>{block.text}</p>
}


const PromptInputWithRef = (props: any) => {
    const defaultValue = useRef("")
    const inputRef = useRef<HTMLParagraphElement>(null)
    const [pointer, setPointer] = useState<number>(-1)
    useEffect(() => {
        if (props.isActive) {
            props.scrollToBottom("instant")
            inputRef.current?.focus()
        }
    }, [props.isActive, props.nfocus])

    const handleInput = (event: any) => {
        if (props.onChange) {
            props.onChange(event.target.textContent)
        }
    }

    const setInputRef = (cmd: string) => {
        if (inputRef.current !== null) {
            inputRef.current.textContent = cmd;
        }
    }
    const handleKeyUp = (event: React.KeyboardEvent) => {
        if (props.confirmInput && event.key.toLocaleLowerCase() === 'enter') {
            props.confirmInput()
            if (inputRef.current) {
                inputRef.current.innerText = ""
            }
            setPointer(-1)
        }
        const updateInputFromPointer = (newPointer: number) => {
            let text = props.history.at(props.history.length - 1 - newPointer)
            if (!text) {
                text = ""
            }
            setInputRef(text)
        }
        if (event.key.toLocaleLowerCase() === 'arrowup') {
            if (pointer !== props.history.length - 1) {
                setPointer(() => (pointer + 1))
                updateInputFromPointer(pointer + 1)
            }
        }
        if (event.key.toLocaleLowerCase() === 'arrowdown') {
            let newPointer
            if (pointer < 0) {
                newPointer = -1
            } else {
                newPointer = pointer - 1
            }
            setPointer(newPointer)
            updateInputFromPointer(newPointer)
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
