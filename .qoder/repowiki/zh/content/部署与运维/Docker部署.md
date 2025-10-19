# Docker部署

<cite>
**Referenced Files in This Document**   
- [Dockerfile](file://server/manifest/docker/Dockerfile)
- [docker.sh](file://server/manifest/docker/docker.sh)
- [entrypoint.sh](file://server/manifest/docker/entrypoint.sh)
- [main.go](file://server/main.go)
- [Makefile](file://server/Makefile)
</cite>

## 目录
1. [Docker镜像构建](#docker镜像构建)
2. [多阶段构建流程](#多阶段构建流程)
3. [容器启动与配置](#容器启动与配置)
4. [启动脚本功能](#启动脚本功能)
5. [最佳实践](#最佳实践)
6. [常见问题排查](#常见问题排查)

## Docker镜像构建

hotgo-2.0后端服务的Docker镜像构建基于标准的Docker工作流程。项目在`server/manifest/docker`目录下提供了完整的Docker部署文件，包括Dockerfile、docker.sh和entrypoint.sh脚本。构建过程首先需要通过Makefile中的`build`目标编译前端资源并生成后端二进制文件，然后使用Docker命令基于Dockerfile构建最终的容器镜像。

**Section sources**
- [Dockerfile](file://server/manifest/docker/Dockerfile#L1-L24)
- [Makefile](file://server/Makefile#L1-L114)

## 多阶段构建流程

hotgo-2.0的Docker部署采用单阶段构建模式，基于`loads/alpine:3.8`基础镜像。构建过程在Dockerfile中定义，首先设置工作目录`/app`，然后将预编译的二进制文件`hotgo`从`./temp/linux_amd64/`目录复制到容器中。同时，将`hack`配置目录、`manifest/config`配置文件、`resource`资源文件以及`entrypoint.sh`启动脚本一并复制到容器内指定位置。最后通过RUN指令为二进制文件和启动脚本添加可执行权限，完成镜像构建。

**Section sources**
- [Dockerfile](file://server/manifest/docker/Dockerfile#L6-L13)

## 容器启动与配置

用户可以通过docker.sh脚本进行容器相关的操作。该脚本在Docker构建前执行，可用于执行预处理任务。容器启动时，通过Dockerfile中的CMD指令指定执行entrypoint.sh启动脚本。环境变量通过ENV指令在构建时设置，主要配置了工作目录`WORKDIR`为`/app`。日志文件和配置文件目录通过ADD指令在构建阶段直接复制到镜像中，确保容器运行时能够访问必要的配置和资源文件。

**Section sources**
- [Dockerfile](file://server/manifest/docker/Dockerfile#L6-L13)
- [docker.sh](file://server/manifest/docker/docker.sh#L1-L9)
- [Makefile](file://server/Makefile#L1-L114)

## 启动脚本功能

entrypoint.sh启动脚本负责容器的最终启动流程。脚本首先切换到`/app`工作目录，然后在后台启动hotgo二进制程序。启动成功后输出"hotgo start all server.."提示信息。为了保持容器运行状态，脚本最后执行`tail -f /dev/null`命令持续输出空设备内容，防止容器因主进程结束而退出。该脚本还负责设置正确的文件权限，确保hotgo二进制文件具有可执行权限。

**Section sources**
- [entrypoint.sh](file://server/manifest/docker/entrypoint.sh#L1-L5)
- [Dockerfile](file://server/manifest/docker/Dockerfile#L12-L13)

## 最佳实践

在Docker部署hotgo-2.0时，建议遵循以下最佳实践：使用Alpine Linux作为基础镜像以减小镜像体积；通过Makefile的`image`目标自动化构建过程，支持版本标签；在构建前确保前端资源已正确编译并放置到指定目录；合理配置容器资源限制；考虑添加健康检查机制以监控服务状态；使用环境变量而非硬编码配置来提高部署灵活性。同时，应定期更新基础镜像以包含最新的安全补丁。

**Section sources**
- [Makefile](file://server/Makefile#L1-L114)
- [Dockerfile](file://server/manifest/docker/Dockerfile#L1-L24)

## 常见问题排查

容器部署过程中可能遇到多种问题。容器启动失败通常源于二进制文件权限不足或依赖文件缺失，可通过检查Dockerfile中的chmod命令和ADD指令确保正确设置。端口冲突问题可通过Docker的端口映射功能解决，确保容器端口与宿主机端口正确映射。文件权限错误常见于挂载的卷，应确保宿主机文件具有适当的读写权限。若服务无法正常启动，可进入容器内部检查日志文件，或通过`docker logs`命令查看启动输出，定位具体错误原因。

**Section sources**
- [Dockerfile](file://server/manifest/docker/Dockerfile#L12-L13)
- [entrypoint.sh](file://server/manifest/docker/entrypoint.sh#L1-L5)
- [Makefile](file://server/Makefile#L1-L114)