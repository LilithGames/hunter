# hunter

> influx时序数据写客户端

    问题：直接使用influx client或API写数据时，没有统一定义，不便统一审查和管理
    解决：用protobuf定义格式规范，protoc生成写入端代码
    
#### 一、格式规范
    * 时序数据表名、标签名和字段名均为snake风格，单词小写且间隔为下划线
    * schema的定义三层嵌套，顶层为类别，中层为表，底层为标签或值
    举例kubernetes数据：
    message kubernetes {
      message node {
        message tag {
          string node_name = 1;
        }
        message field {
          int64 cpu_usage_core_nanoseconds = 1;
          int64 memory_available_bytes = 2;
        }
      }
      message pod_network {
        message tag {
          string namespace = 1;
          string node_name = 2;
          string pod_name = 3;
        }
        message field {
          int64 rx_bytes = 1;
          int64 rx_errors = 2;
        }
      }
      message pod_volume {
        message tag {
          string namespace = 1;
          string node_name = 2;
          string pod_name = 3;
          string volume_name = 4;
        }
        message field {
          int64 available_bytes = 1;
          int64 capacity_bytes = 2;
          int64 used_bytes = 3;
        }
      }
      message pod_container {
        message tag {
          string node_name = 1;
          string container_name = 2;
        }
        message field {
          int64 cpu_usage_core_nanoseconds = 1;
          int64 memory_usage_bytes = 2;
        }
      }
    }
 
 #### 二、代码生成

    安装：
    * go install github.com/LilithGames/hunter/protoc-gen-series
    * go install github.com/golang/protobuf/protoc-gen-go
    使用：
    1. 编写指定时序数据的protobuf文件
    2. 使用protoc生成代码，如protoc -I . --go_out=. --series_out=. ./testdata/proto/*