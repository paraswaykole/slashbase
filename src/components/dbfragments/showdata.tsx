import styles from './showdata.module.scss'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { DBConnection, DBDataModel, DBQueryData } from '../../data/models'
import apiService from '../../network/apiService'
import { selectDBConnection, selectDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppSelector } from '../../redux/hooks'
import Table from './table/table'

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
    const [dataLoading, setDataLoading] = useState(false)
    
    useEffect(()=>{
        const dModel = dbDataModels.find(x => x.schemaName == mschema && x.name == mname)
        if(dModel){
            setDataModel(dModel)
        }
        // else redirect to home fragment             
    }, [dbDataModels])

    useEffect(()=>{
        if (dataModel && !queryCount){ 
            fetchData(true)
        }
    }, [dataModel, queryCount])
    
    useEffect(()=>{
        fetchData(false)
    }, [queryOffset])
    
    const fetchData = async (fetchCount: boolean) => {
        if(!dataModel || dataLoading) return
        setDataLoading(true)
        const result = await apiService.getDBDataInDataModel(dbConnection!.id, dataModel!.schemaName ?? '', dataModel!.name, queryOffset, fetchCount)
        if (result.success) {
            setQueryModel(result.data)
            if (!queryCount){
                setQueryCount(result.data.count)
            }
        }
        setDataLoading(false)
    }

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
            { queryData && 
                <div className={styles.tableContainer}>
                    <Table queryData={queryData}/>
                </div> 
            }
            <br/><br/><br/>
            <div className={styles.bottomBar}>
                {dataLoading ? 
                    <progress className="progress is-primary" max="100">loading</progress>
                    :
                    <nav className="pagination is-centered is-rounded" role="navigation" aria-label="pagination">
                        <button className="button pagination-previous" onClick={onPreviousPage}>Previous</button>
                        <button className="button pagination-next" onClick={onNextPage}>Next</button>
                        <ul className="pagination-list">
                            Showing {queryOffset} - {queryOffsetRangeEnd} of {queryCount}
                        </ul>
                    </nav>
                }
            </div>
        </React.Fragment>
    )
}


export default DBShowDataFragment