import {
    Suspense,
    useContext,
    useEffect,
    useState,
    FC,
    MouseEvent
} from 'react'

import { ConfigProvider, Layout, Space, Watermark } from 'antd'
import { Outlet, useLocation } from 'react-router-dom'
import { CopyrightOutlined } from '@ant-design/icons'

import { GlobalContext } from '@/context'
import Loading from '../Loading'
import HeaderTitle from './HeaderTitle/HeaderTitle'
import { SiderMenu } from './SiderMenu/SiderMenu'
import { UserInfo } from './UserInfo/UserInfo'
import { Msg } from './Msg/Msg'
import { Setting } from './Setting/Setting'
import { SpaceInfo } from './SpaceInfo/SpaceInfo'

import styles from './style/index.module.less'

const { Header, Footer, Sider, Content } = Layout

export type PromLayoutProps = {
    watermark?: string
}

export const LayoutContentID = 'LayoutContent'

const PromLayout: FC<PromLayoutProps> = (props) => {
    const { user, setLayoutContentElement } = useContext(GlobalContext)

    const { watermark = user?.username } = props
    const [collapsed, setCollapsed] = useState(true)

    const handleOnMouseEnter = (e: MouseEvent) => {
        // 鼠标移入, 展开菜单
        if (e.type === 'mouseenter') {
            setCollapsed(false)
        }
        // 鼠标移出, 收起菜单
        if (e.type === 'mouseleave') {
            setCollapsed(true)
        }
    }

    useEffect(() => {
        setLayoutContentElement?.(document.getElementById(LayoutContentID))
    }, [])

    const local = useLocation()

    useEffect(() => {
        // TODO 做路由权限认证
        console.log('local.pathname', local)
    }, [local.pathname])

    return (
        <Layout className={[styles.widthHight100, styles.Layout].join(' ')}>
            <ConfigProvider
                theme={{
                    token: {},
                    components: {
                        Layout: {
                            colorBgHeader: '#FFF',
                            colorBgBody: '#f0f0f0'
                        },
                        Menu: {
                            colorBgContainer: '#fafafa',
                            itemBorderRadius: 8
                        }
                    }
                }}
            >
                <Header className={styles.LayoutHeader}>
                    <HeaderTitle />
                    <Space size={12} direction="horizontal">
                        <SpaceInfo />
                        <Msg />
                        <Setting />
                        <UserInfo />
                    </Space>
                </Header>

                <Layout>
                    <Sider
                        defaultCollapsed
                        collapsed={collapsed}
                        className={styles.LayoutSider}
                        onMouseEnter={handleOnMouseEnter}
                        onMouseLeave={handleOnMouseEnter}
                        collapsedWidth={60}
                    >
                        <SiderMenu />
                    </Sider>

                    <Layout>
                        <Content>
                            <Suspense fallback={<Loading />}>
                                <div
                                    className={styles.LayoutContent}
                                    id={LayoutContentID}
                                >
                                    <Watermark
                                        content={watermark}
                                        className="wh100"
                                    >
                                        <Outlet />
                                    </Watermark>
                                </div>
                            </Suspense>
                        </Content>
                        <Footer className={styles.LayoutFooter}>
                            <CopyrightOutlined />
                            {window.location.host}
                        </Footer>
                    </Layout>
                </Layout>
            </ConfigProvider>
        </Layout>
    )
}

export default PromLayout
