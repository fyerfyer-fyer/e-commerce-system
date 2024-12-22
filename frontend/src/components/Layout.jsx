import React from 'react';
import { Layout } from 'antd';
import AppHeader from './Header';
import AppFooter from './Footer';
import AppSidebar from './Sidebar';

const { Content } = Layout;

function AppLayout({ children }) {
  return (
    <Layout>
      <AppHeader />
      <Layout>
        <AppSidebar />
        <Content style={{ padding: '20px', background: '#fff', minHeight: 'calc(100vh - 120px)' }}>
          {children}
        </Content>
      </Layout>
      <AppFooter />
    </Layout>
  );
}

export default AppLayout;