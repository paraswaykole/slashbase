import React, { FunctionComponent, useEffect, useState } from 'react'

import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import DBShowDataFragment from '../../components/dbfragments/showdata'
import { DBConnection } from '../../data/models'
import Constants from '../../constants'
import DBShowModelFragment from '../../components/dbfragments/showmodel'
import { Link, useParams, useSearchParams } from 'react-router-dom'

const DBPathPage: FunctionComponent<{}> = () => {

    const { id, path } = useParams()
    const [searchParams] = useSearchParams()
    const mschema = searchParams.get("mschema")
    const mname = searchParams.get("mname")

    const [error404, setError404] = useState(false)
    const dispatch = useAppDispatch()

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    useEffect(() => {
        (async () => {
            if (id) {
                try {
                    await dispatch((getDBConnection({ dbConnId: String(id) }))).unwrap()
                } catch (e) {
                    setError404(true)
                    return
                }
                dispatch((getDBDataModels({ dbConnId: String(id) })))
                dispatch((getDBQueries({ dbConnId: String(id) })))
            }
        })()
    }, [dispatch, id])

    if (error404 || !['data', 'model'].includes(String(path))) {
        return (<h1>404 not found</h1>)
    }

    return (
        <React.Fragment>
            {mname &&
                <React.Fragment>
                    <div className="tabs is-toggle is-toggle-rounded tabs-set">
                        <h1>{`Showing ${mschema === '' ? mname : `${mschema}.${mname}`}`}</h1>
                        <ul>
                            <li className={path === 'data' ? 'is-active' : ''}>
                                <Link to={Constants.APP_PATHS.DB_PATH.path.replace('[id]', String(id)).replace('[path]', String('data')) + "?mschema=" + mschema + "&mname=" + mname}>
                                    <span className="icon is-small"><i className="fas fa-table" aria-hidden="true" /></span>
                                    <span>Data</span>
                                </Link>
                            </li>
                            <li className={path === 'model' ? 'is-active' : ''}>
                                <Link to={Constants.APP_PATHS.DB_PATH.path.replace('[id]', String(id)).replace('[path]', String('model')) + "?mschema=" + mschema + "&mname=" + mname}>
                                    <span className="icon is-small"><i className="fas fa-list-alt" aria-hidden="true" /></span>
                                    <span>Model</span>
                                </Link>
                            </li>
                        </ul>
                    </div>
                    {path === 'data' && <DBShowDataFragment key={String(mschema) + String(mname)} />}
                    {path === 'model' && <DBShowModelFragment />}
                </React.Fragment>
            }
            <style>{`
                .tabs-set {
                    align-items: center;
                    justify-content: space-between;
                    margin-bottom: 0.7rem;
                }
                .tabs-set ul{
                    flex-grow: 0;
                }
            `}</style>
        </React.Fragment>
    )
}

export default DBPathPage
