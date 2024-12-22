import React from 'react';
import { Card, Table, Typography, Space, Modal, Button } from 'antd';
import { Link } from 'react-router-dom';

const { Title } = Typography;

const data = [
  {
    key: '1',
    reply_id: 'R001',
    content: 'This is a great comment!',
    user_name: 'John Doe',
  },
  {
    key: '2',
    reply_id: 'R002',
    content: 'I agree with the previous comment.',
    user_name: 'Jane Smith',
  },
];

function ReplyListPage({ match }) {
  const commentId = match.params.commentId; // Assuming you're using react-router v5

  // 定义删除确认弹窗函数
  const showDeleteConfirm = (record) => {
    Modal.confirm({
      title: 'Delete Reply',
      content: `Are you sure you want to delete this reply?`,
      okText: 'Delete',
      okType: 'danger',
      cancelText: 'Cancel',
      onOk() {
        console.log('Deleted reply:', record);
        // 在这里添加删除回复的后端逻辑
      },
    });
  };

  // 将 columns 移动到组件内部
  const columns = [
    {
      title: 'Reply ID',
      dataIndex: 'reply_id',
      key: 'reply_id',
    },
    {
      title: 'Content',
      dataIndex: 'content',
      key: 'content',
    },
    {
      title: 'User',
      dataIndex: 'user_name',
      key: 'user_name',
    },
    {
      title: 'Actions',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <a href={`/reply/edit/${record.reply_id}`}>Edit</a>
          <a href="#" onClick={() => showDeleteConfirm(record)}>Delete</a>
        </Space>
      ),
    },
  ];

  return (
    <div className="page-container">
      <Title level={2}>Replies for Comment {commentId}</Title>
      <Card title="Reply List" bordered={false}>
        <Table columns={columns} dataSource={data} pagination={{ pageSize: 5 }} />
      </Card>
      <Link to="/reply/add" style={{ display: 'block', marginTop: 16 }}>
        <Button type="primary">Add New Reply</Button>
      </Link>
    </div>
  );
}

export default ReplyListPage;