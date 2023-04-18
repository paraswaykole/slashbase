import styles from './cheatsheet.module.scss'
import React from 'react'

type CheatsheetCommandPropType = {
    cmd: { title: string, description: string, command: string },
    isLast: boolean
}

const CheatsheetCommand = ({ cmd, isLast }: CheatsheetCommandPropType) => {

    return (<React.Fragment>
        <div>
            <h2>{cmd.title}</h2>
            <p>{cmd.description}</p>
            <code>{cmd.command}</code>
            {!isLast && <hr />}
        </div>

    </React.Fragment>)
}


export default CheatsheetCommand