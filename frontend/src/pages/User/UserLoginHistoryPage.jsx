import React from 'react';
import { Card, Table, Typography } from 'antd';

const { Title } = Typography;

const columns = [
  {
    title: 'Login Time',
    dataIndex: 'login_time',
    key: 'login_time',
  },
  {
    title: 'IP Address',
    dataIndex: 'login_ip',
    key: 'login_ip',
  },
];

const data = [
  {
    key: '1',
    login_time: '2023-10-01 10:00:00',
    login_ip: '192.168.1.1',
  },
  {
    key: '2',
    login_time: '2023-10-02 12:30:00',
    login_ip: '192.168.1.2',
  },
];

function UserLoginHistoryPage() {
  return (
    <div className="page-container">
      <Title level={2}>Login History</Title>
      <Card title="User Login History" bordered={false}>
        <Table columns={columns} dataSource={data} pagination={{ pageSize: 5 }} />
      </Card>
    </div>
  );
}

export default UserLoginHistoryPage;