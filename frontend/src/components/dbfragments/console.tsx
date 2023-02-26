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

    useEffect(() => {
        dispatch(initConsole(dbConnection!.id))
    }, [dbConnection])

    const confirmInput = () => {
        dispatch(runConsoleCmd({ dbConnId: dbConnection!.id, cmdString: input }))
        setInput('')
    }

    useEffect(() => {
        consoleEndRef.current?.scrollIntoView({ behavior: 'smooth' })
    }, [output])

    return <div className={currentTab.isActive ? "db-tab-active" : "db-tab"}>
        <div className={styles.console}>
            {output.map(block => {
                return <OutputBlock block={block} />
            })}
            <PromptInputWithRef onChange={setInput} confirmInput={confirmInput} />
            <span ref={consoleEndRef}></span>
        </div>
    </div>
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
