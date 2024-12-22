import React from 'react';
import { Card, Table, Typography, Space, Tag } from 'antd';
import { Link } from 'react-router-dom';

const { Title } = Typography;

const columns = [
  {
    title: 'Event ID',
    dataIndex: 'event_id',
    key: 'event_id',
  },
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: 'Start Time',
    dataIndex: 'start_time',
    key: 'start_time',
  },
  {
    title: 'End Time',
    dataIndex: 'end_time',
    key: 'end_time',
  },
  {
    title: 'Status',
    dataIndex: 'status',
    key: 'status',
    render: (text) => (
      <Tag color={text === 'upcoming' ? 'blue' : text === 'ongoing' ? 'green' : 'red'}>{text}</Tag>
    ),
  },
  {
    title: 'Actions',
    key: 'action',
    render: (_, record) => (
      <Space size="middle">
        <Link to={`/seckill/products?eventId=${record.event_id}`}>View Products</Link>
      </Space>
    ),
  },
];

const data = [
  {
    key: '1',
    event_id: 'E001',
    name: 'Summer Sale',
    start_time: '2023-10-01 10:00:00',
    end_time: '2023-10-07 23:59:59',
    status: 'upcoming',
  },
  {
    key: '2',
    event_id: 'E002',
    name: 'Black Friday',
    start_time: '2023-11-24 00:00:00',
    end_time: '2023-11-24 23:59:59',
    status: 'ongoing',
  },
];

function SeckillEventsPage() {
  return (
    <div className="page-container">
      <Title level={2}>Seckill Events</Title>
      <Card title="Seckill Event List" bordered={false}>
        <Table columns={columns} dataSource={data} pagination={{ pageSize: 5 }} />
      </Card>
    </div>
  );
}

export default SeckillEventsPage;