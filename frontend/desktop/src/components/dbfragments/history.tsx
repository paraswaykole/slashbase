import React, { useContext, useEffect } from 'react'
import { DBConnection, Tab } from '../../data/models'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import toast from 'react-hot-toast'
import InfiniteScroll from 'react-infinite-scroll-component'
import dateformat from 'dateformat'
import { getDBQueryLogs, reset, selectDBQueryLogs, selectDBQueryLogsNext } from '../../redux/dbHistorySlice'
import TabContext from '../layouts/tabcontext'
import Button from '../ui/Button'


type DBHistoryPropType = {
}

const DBHistoryFragment = ({ }: DBHistoryPropType) => {

    const dispatch = useAppDispatch()

    const currentTab: Tab = useContext(TabContext)!

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)
    const dbQueryLogs = useAppSelector(selectDBQueryLogs)
    const dbQueryLogsNext = useAppSelector(selectDBQueryLogsNext)

    useEffect(() => {
        if (dbConnection) {
            (async () => {
                dispatch(reset())
            })()
            fetchDBQueryLogs()
        }
    }, [dispatch, dbConnection])

    const fetchDBQueryLogs = async () => {
        const result = await dispatch(getDBQueryLogs({ dbConnId: dbConnection!.id })).unwrap()
        if (!result.success) {
            toast.error(result.error!)
        }
    }

    function refreshHandler() {
        dispatch(reset())
        fetchDBQueryLogs()
    }

    return (
        <div className={currentTab.isActive ? "db-tab-active" : "db-tab"}>
            {dbConnection &&
                <React.Fragment>
                    <div className="is-flex is-justify-content-space-between">
                        <h1>Showing History in {dbConnection.name}</h1>
                        <Button
                            text='Refresh' 
                            icon={<i className="fas fa-sync"/>}
                            onClick={refreshHandler}
                        />
                    </div>
                    <br />
                    <InfiniteScroll
                        dataLength={dbQueryLogs.length}
                        next={fetchDBQueryLogs}
                        hasMore={dbQueryLogsNext !== -1}
                        loader={
                            <p style={{ textAlign: 'center' }}>
                                Loading...
                            </p>
                        }
                        endMessage={
                            <p style={{ textAlign: 'center' }}>
                                <b>You have seen it all!</b>
                            </p>
                        }
                        scrollableTarget="maincontent"
                    >
                        <table className={"table is-bordered is-striped is-narrow is-hoverable is-fullwidth"}>
                            <tbody>
                                {dbQueryLogs.map((log) => (
                                    <tr key={log.id}>
                                        <td>
                                            <code>{log.query}</code>
                                        </td>
                                        <td style={{ fontSize: '14px', width: '120px' }}>
                                            {dateformat(log.createdAt, "mmm dd, yyyy HH:MM:ss")}
                                        </td>
                                    </tr>
                                )
                                )}
                            </tbody>
                        </table>
                    </InfiniteScroll>
                </React.Fragment>
            }
        </div>
    )
}


export default DBHistoryFragment