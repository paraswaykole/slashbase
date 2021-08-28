import styles from './table.module.scss'
import React from 'react'
import { DBQueryData } from '../../data/models'


type ProjectCardPropType = { 
    queryData: DBQueryData
}

const Table = ({queryData}: ProjectCardPropType) => {

    return (
        <React.Fragment>
            <table className={"table is-bordered is-striped is-narrow is-hoverable is-fullwidth"}>
                <thead>
                    <tr>
                        {queryData.columns.map(colName => (<th key={colName}>{colName}</th>))}    
                    </tr>
                </thead>
                <tbody>
                    {queryData?.rows.map((row,index)=> {
                        return <tr key={index}>
                            { queryData.columns.map((colName, index) => {
                                    return <td key={colName+index}>{row[colName] ? row[colName] : <span className={styles.nullValue}>NULL</span>}</td>
                                })
                            }
                        </tr>
                    })}
                </tbody>
            </table>
        </React.Fragment>
    )
}


export default Table