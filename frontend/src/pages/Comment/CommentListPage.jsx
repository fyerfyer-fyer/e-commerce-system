import React from 'react';
import { Card, Table, Typography, Rate, Space, Button, Modal } from 'antd';
import { Link } from 'react-router-dom';

const { Title } = Typography;

const data = [
  {
    key: '1',
    comment_id: 'C001',
    product_id: 'P001',
    rating: 5,
    content: 'This is an excellent product!',
  },
  {
    key: '2',
    comment_id: 'C002',
    product_id: 'P002',
    rating: 4,
    content: 'Good quality, but could be better.',
  },
];

function CommentListPage({ match }) {
  const productId = match.params.productId; // Assuming you're using react-router v5

  // 定义删除确认弹窗函数
  function showDeleteConfirm(record) {
    Modal.confirm({
      title: 'Delete Comment',
      content: `Are you sure you want to delete this comment?`,
      okText: 'Delete',
      okType: 'danger',
      cancelText: 'Cancel',
      onOk() {
        console.log('Deleted comment:', record);
        // 在这里添加删除评论的逻辑，例如调用后端 API
      },
    });
  }

  // 将 columns 移动到组件内部
  const columns = [
    {
      title: 'Product ID',
      dataIndex: 'product_id',
      key: 'product_id',
    },
    {
      title: 'Rating',
      dataIndex: 'rating',
      key: 'rating',
      render: (text) => <Rate disabled defaultValue={text} />,
    },
    {
      title: 'Comment',
      dataIndex: 'content',
      key: 'content',
    },
    {
      title: 'Actions',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Link to={`/comment/edit/${record.comment_id}`}>Edit</Link>
          <a href="#" onClick={() => showDeleteConfirm(record)}>Delete</a>
        </Space>
      ),
    },
  ];

  return (
    <div className="page-container">
      <Title level={2}>Comments for Product {productId}</Title>
      <Card title="Comment List" bordered={false}>
        <Table columns={columns} dataSource={data} pagination={{ pageSize: 5 }} />
      </Card>
      <Link to="/comment/add" style={{ display: 'block', marginTop: 16 }}>
        <Button type="primary">Add New Comment</Button>
      </Link>
    </div>
  );
}

export default CommentListPage;