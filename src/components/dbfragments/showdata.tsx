import styles from './showdata.module.scss'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { DBConnection, DBDataModel, DBQueryData } from '../../data/models'
import apiService from '../../network/apiService'
import { selectDBConnection, selectDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'

type DBShowDataPropType = { 

}

const DBShowDataFragment = (_: DBShowDataPropType) => {

    const router = useRouter()
    const { mschema, mname } = router.query

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbDataModels: DBDataModel[] = useAppSelector(selectDBDataModels)
    
    const [dataModel, setDataModel] = useState<DBDataModel>()
    const [queryData, setQueryModel] = useState<DBQueryData>()
    const [queryOffset, setQueryOffset] = useState(0)
    const [queryCount, setQueryCount] = useState<number|undefined>(undefined)
    const [queryLimit] = useState(200)
    
    useEffect(()=>{
        const dModel = dbDataModels.find(x => x.schemaName == mschema && x.name == mname)
        if(dModel){
            setDataModel(dModel)
            {(async () => {
                const result = await apiService.getDBDataInDataModel(dbConnection!.id, dModel!.schemaName ?? '', dModel!.name, queryOffset, !queryCount)
                if (result.success) {
                    setQueryModel(result.data)
                    if (!queryCount){
                        setQueryCount(result.data.count)
                    }
                }
            })()}
        }
        // else redirect to home fragment             
    }, [dbConnection, dbDataModels, queryOffset, queryCount])

    const onPreviousPage = () => {
        let previousOffset = queryOffset - queryLimit
        if(previousOffset < 0) {
            previousOffset = 0
        }
        setQueryOffset(previousOffset)
    }
    const onNextPage = () => {
        let nextOffset = queryOffset + queryLimit
        if (nextOffset > (queryCount ?? 0)){
            return
        }
        setQueryOffset(nextOffset)
    }

    const queryOffsetRangeEnd = (queryData?.rows.length ?? 0) === queryLimit ? 
        queryOffset + queryLimit : queryOffset + (queryData?.rows.length ?? 0)

    return (
        <React.Fragment>
            <h1>Showing {dataModel?.schemaName}.{dataModel?.name}</h1>
            <table className={"table "+styles.tableContainer}>
            <thead>
                <tr>
                    {queryData?.columns.map(colName => (<th key={colName}>{colName}</th>))}    
                </tr>
            </thead>
            <tbody>
                {queryData?.rows.map((row,index)=> {
                    return <tr key={index}>
                        { queryData?.columns.map((colName, index) => {
                                return <td key={colName+index}>{row[colName]}</td>
                            })
                        }
                    </tr>
                })}
                
            </tbody>
            </table>
            <nav className="pagination is-centered" role="navigation" aria-label="pagination">
                <a className="pagination-previous" onClick={onPreviousPage}>Previous</a>
                <a className="pagination-next" onClick={onNextPage}>Next</a>
                <ul className="pagination-list">
                    Showing {queryOffset} - {queryOffsetRangeEnd} of {queryCount}
                </ul>
            </nav>
        </React.Fragment>
    )
}


export default DBShowDataFragment