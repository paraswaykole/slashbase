import styles from './chart.module.scss'
import React, { useRef, useState } from 'react'
import { DBConnection, DBQueryData } from '../../../data/models'
import { DBConnType } from '../../../data/defaults'
import { Bar, Line, Pie } from 'react-chartjs-2'
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    LineElement,
    PointElement,
    ArcElement,
    Tooltip
} from 'chart.js'

ChartJS.register(
    CategoryScale,
    LinearScale,
    BarElement,
    LineElement,
    PointElement,
    ArcElement,
    Tooltip
)


type ChartPropType = {
    dbConn: DBConnection
    queryData: DBQueryData,
}

type ChartDataType = {
    xaxis: string,
    yaxis: string,
    data: {
        labels: string[]
        datasets: Array<{
            label: string,
            data: number[],
            backgroundColor: string
        }>
    }
}

enum ChartViewType {
    NONE = "NONE", // default
    BARCHART = "BARCHART",
    LINECHART = "LINECHART",
    PIECHART = "PIECHART"
}

const Chart = ({ dbConn, queryData }: ChartPropType) => {

    const [chartViewType, setChartViewType] = useState<ChartViewType>(ChartViewType.NONE)
    const [chartData, setChartData] = useState<ChartDataType>()

    const selectChartTypeRef = useRef<HTMLSelectElement>(null)
    const selectXAxisRef = useRef<HTMLSelectElement>(null)
    const selectYAxisRef = useRef<HTMLSelectElement>(null)

    const keys = (dbConn.type === DBConnType.POSTGRES || dbConn.type === DBConnType.MYSQL) ? queryData.columns : queryData.keys

    const createChart = () => {
        const xaxis = selectXAxisRef.current!.value
        const yaxis = selectYAxisRef.current!.value
        let labels: string[]
        let data: number[]
        if (dbConn.type === DBConnType.POSTGRES || dbConn.type === DBConnType.MYSQL) {
            const xColIdx = keys.findIndex(x => x === xaxis)
            const yColIdx = keys.findIndex(y => y === yaxis)
            labels = queryData.rows.map(row => row[xColIdx])
            data = queryData.rows.map(row => row[yColIdx])
        } else {
            labels = queryData.data.map(d => d[xaxis])
            data = queryData.data.map(d => d[yaxis])
        }
        setChartViewType(ChartViewType[selectChartTypeRef.current!.value.toString() as keyof typeof ChartViewType])
        setChartData({
            xaxis: xaxis,
            yaxis: yaxis,
            data: {
                labels,
                datasets: [{
                    label: yaxis,
                    data: data,
                    backgroundColor: '#615f9c',
                }]
            }
        })
    }

    return <React.Fragment>
        <div className="card">
            <div className="card-content">
                <div className={"content " + styles.contentCenter}>
                    {chartViewType === ChartViewType.NONE &&
                        <React.Fragment>
                            <h2 className="title is-2">Create chart from the data</h2>
                            <div className="field">
                                <label className="label">Select chart type:</label>
                                <div className="control">
                                    <div className="select">
                                        <select ref={selectChartTypeRef}>
                                            <option value={ChartViewType.BARCHART}>Bar Chart</option>
                                            <option value={ChartViewType.LINECHART}>Line Chart</option>
                                            <option value={ChartViewType.PIECHART}>Pie Chart</option>
                                        </select>
                                    </div>
                                </div>
                            </div>
                            <div className="field">
                                <label className="label">Select x-axis:</label>
                                <div className="control">
                                    <div className="select">
                                        <select ref={selectXAxisRef}>
                                            {keys.map(key => (
                                                <option key={key} value={key}>{key}</option>
                                            ))}
                                        </select>
                                    </div>
                                </div>
                            </div>
                            <div className="field">
                                <label className="label">Select y-axis:</label>
                                <div className="control">
                                    <div className="select">
                                        <select ref={selectYAxisRef}>
                                            {keys.map(key => (
                                                <option key={key} value={key}>{key}</option>
                                            ))}
                                        </select>
                                    </div>
                                </div>
                            </div>
                            <br />
                            <div className="control">
                                <button className="button is-primary" onClick={createChart}>Create</button>
                            </div>

                        </React.Fragment>
                    }
                    {chartViewType === ChartViewType.BARCHART &&
                        <React.Fragment>
                            <h2 className="title is-2">Bar chart</h2>
                            <div className={styles.barChartWrapper}>
                                <Bar data={chartData!.data} />
                            </div>
                        </React.Fragment>
                    }
                    {chartViewType === ChartViewType.LINECHART &&
                        <React.Fragment>
                            <h2 className="title is-2">Line chart</h2>
                            <div className={styles.lineChartWrapper}>
                                <Line data={chartData!.data} />
                            </div>
                        </React.Fragment>
                    }
                    {chartViewType === ChartViewType.PIECHART &&
                        <React.Fragment>
                            <h2 className="title is-2">Pie chart</h2>
                            <div className={styles.pieChartWrapper}>
                                <Pie data={chartData!.data} />
                            </div>
                        </React.Fragment>
                    }
                </div>
            </div>
        </div>
    </React.Fragment>
}

export default Chart