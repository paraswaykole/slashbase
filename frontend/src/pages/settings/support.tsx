import type { NextPage } from 'next'
import Link from 'next/link'
import React, { useEffect, useRef, useState } from 'react'
import AppLayout from '../../components/layouts/applayout'
import Constants from '../../constants'

const SupportPage: NextPage = () => {

    return (
        <AppLayout title="Support | Slashbase">

            <h1>Support</h1>

            <p>We are more than happy to help you.</p>
            <br />
            <h4>Have questions/ideas or need help?</h4>
            <p>If you are have any questions or ideas,
                please join our Discord server and share. We are open to discussing anything your questions or ideas.</p>
            <Link href={Constants.EXTERNAL_PATHS.DISCORD_COMMUNITY}>
                <a target="_blank">
                    <button className="button is-secondary">
                        <i className={"fab fa-discord"} />&nbsp;&nbsp;Join Discord Community
                    </button>
                </a>
            </Link>
            <br /><br />
            <h4>Facing bugs or errors?</h4>
            <p>If you are getting any errors or bugs, join our Discord server and ask for help.
                We will try to fix the errors bugs in the next version release.</p>
            <Link href={Constants.EXTERNAL_PATHS.REPORT_BUGS}>
                <a target="_blank">
                    <button className="button is-secondary">
                        <i className={"fab fa-github"} />&nbsp;&nbsp;File a GitHub Issue.
                    </button>
                </a>
            </Link>
            <br /><br />
            <h4>Want to check what&apos;s new?</h4>
            <p>Vist our website to check new releases.</p>
            <Link href={Constants.EXTERNAL_PATHS.OFFICIAL_WEBSITE}>
                <a target="_blank">
                    <button className="button is-secondary">
                        <i className={"fas fa-globe"} />&nbsp;&nbsp;Visit Slashbase.com
                    </button>
                </a>
            </Link>
        </AppLayout>
    )
}

export default SupportPage
