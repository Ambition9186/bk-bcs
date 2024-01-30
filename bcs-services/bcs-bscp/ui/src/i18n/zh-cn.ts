export default {
  // 服务列表
  我的服务: '我的服务',
  全部服务: '全部服务',
  新建服务: '新建服务',
  你尚未创建或加入任何服务: '你尚未创建或加入任何服务',
  立即创建: '立即创建',
  申请权限: '申请权限',
  服务名称: '服务名称',
  form_服务名称: '服务名称',
  服务属性: '服务属性',
  服务别名: '服务别名',
  form_服务别名: '服务别名',
  数据格式: '数据格式',
  保存: '保存',
  编辑服务: '编辑服务',
  数据类型: '数据类型',
  文件型: '文件型',
  键值型: '键值型',
  任意类型: '任意类型',
  配置管理: '配置管理',
  所属业务: '所属业务',
  服务描述: '服务描述',
  提交: '提交',
  取消: '取消',
  接入方式: '接入方式',
  创建者: '创建者',
  创建时间: '创建时间',
  配置文件: '配置文件',
  配置项: '配置项',
  个配置文件: '个配置文件',
  个配置项: '个配置项',
  未更新: '未更新',
  最新上线: '最新上线',
  申请服务权限: '申请服务权限',
  服务新建成功: '服务新建成功',
  接下来你可以在服务下新增配置文件: '接下来你可以在服务下新增配置文件',
  接下来你可以在服务下新增配置项: '接下来你可以在服务下新增配置项',
  新增配置文件: '新增配置文件',
  新增配置项: '新增配置项',
  稍后再说: '稍后再说',
  编辑: '编辑',
  最小长度2个字符: '最小长度2个字符',
  最大长度32个字符: '最大长度32个字符',
  '服务名称由英文、数字、下划线、中划线组成且以英文、数字开头和结尾': '服务名称由英文、数字、下划线、中划线组成且以英文、数字开头和结尾',
  '服务别名由中文、英文、数字、下划线、中划线且必须以中文、英文、数字开头和结尾': '服务别名由中文、英文、数字、下划线、中划线且必须以中文、英文、数字开头和结尾',
  '请输入2-32字符，只允许英文、数字、下划线、中划线且必须以英文、数字开头和结尾': '请输入2-32字符，只允许英文、数字、下划线、中划线且必须以英文、数字开头和结尾',
  '请输入2-128字符，只允许中文、英文、数字、下划线、中划线且必须以中文、英文、数字开头和结尾': '请输入2-128字符，只允许中文、英文、数字、下划线、中划线且必须以中文、英文、数字开头和结尾',
  服务描述限制200字符: '服务描述限制200字符',
  更新者: '更新者',
  tips: {
    config: `文件型：通常以文件的形式存储,通常具有良好的可读性和可维护性
    键值型：以键值对的形式存储，其中键（key）用于位置标识一个配置项，值（value）为该配置项的具体内容，kv型配置通常存储在数据库，使用SDK或API的方式读取`,
    type: `任意类型：不对配置项的类型做限制。如果选择下方某个类型，则只能创建指定类型的配置项
         string：单行字符串
         number：数值，包含整数、浮点数、会校验数据类型
         text：多行字符串文本，不校验数据结构，大小2Mb
         json、xml、yaml：不同格式的结构化数据，会校验数据结构`,
  },

  // 导航栏
  服务配置中心: '服务配置中心',
  服务管理: '服务管理',
  分组管理: '分组管理',
  全局变量: '全局变量',
  配置模板: '配置模板',
  脚本管理: '脚本管理',
  服务密钥: '服务密钥',
  产品文档: '产品文档',
  版本日志: '版本日志',
  功能特性: '功能特性',
  问题反馈: '问题反馈',
  退出登录: '退出登录',
  新建业务: '新建业务',

  // 配置管理
  版本名称: '版本名称',
  未命名版本: '未命名版本',
  版本对比: '版本对比',
  版本废弃: '版本废弃',
  确认废弃该版本: '确认废弃该版本',
  '此操作不会删除版本，如需找回或彻底删除请去版本详情的废弃版本列表操作': '此操作不会删除版本，如需找回或彻底删除请去版本详情的废弃版本列表操作',
  可用版本: '可用版本',
  废弃版本: '废弃版本',
  '版本名称/版本说明/修改人': '版本名称/版本说明/修改人',
  版本: '版本',
  版本描述: '版本描述',
  已上线分组: '已上线分组',
  创建人: '创建人',
  生成时间: '生成时间',
  状态: '状态',
  已废弃: '已废弃',
  未上线: '未上线',
  已上线: '已上线',
  操作: '操作',
  只支持未上线版本: '只支持未上线版本',
  恢复: '恢复',
  删除: '删除',
  版本详情: '版本详情',
  '配置文件名/创建人/修改人': '配置文件名/创建人/修改人',
  新建配置文件: '新建配置文件',
  新建配置项: '新建配置项',
  手动新增: '手动新增',
  批量上传: '批量上传',
  从配置模板导入: '从配置模板导入',
  批量导入: '批量导入',
  配置文件名称: '配置文件名称',
  '请输入1~64个字符，只允许英文、数字、下划线、中划线或点': '请输入1~64个字符，只允许英文、数字、下划线、中划线或点',
  配置文件路径: '配置文件路径',
  '客户端拉取配置文件后存放路径为：临时目录/业务ID/服务名称/files/配置文件路径，除了配置文件路径其它参数都在客户端sidecar中配置': '客户端拉取配置文件后存放路径为：临时目录/业务ID/服务名称/files/配置文件路径，除了配置文件路径其它参数都在客户端sidecar中配置',
  请输入绝对路径: '请输入绝对路径',
  配置文件描述: '配置文件描述',
  配置文件格式: '配置文件格式',
  文件权限: '文件权限',
  请输入三位权限数字: '请输入三位权限数字',
  '只能输入三位 0~7 数字': '只能输入三位 0~7 数字',
  我知道了: '我知道了',
  读: '读',
  写: '写',
  执行: '执行',
  用户: '用户',
  用户组: '用户组',
  配置内容: '配置内容',
  文件大小100M以内: '文件大小100M以内',
  '属主（own）': '属主（own）',
  '属组（group）': '属组（group）',
  '其他人（other）': '其他人（other）',
  最大长度64个字符: '最大长度64个字符',
  '请使用中文、英文、数字、下划线、中划线或点': '请使用中文、英文、数字、下划线、中划线或点',
  '文件权限 不能为空': '文件权限 不能为空',
  文件own必须有读取权限: '文件own必须有读取权限',
  最大长度1024个字符: '最大长度1024个字符',
  '无效的路径,路径不符合Unix文件路径格式规范': '无效的路径,路径不符合Unix文件路径格式规范',
  最大长度200个字符: '最大长度200个字符',
  请上传文件: '请上传文件',
  配置内容不能超过50M: '配置内容不能超过50M',
  配置项名称: '配置项名称',
  配置项值: '配置项值',
  '请输入(仅支持大小不超过2M)': '请输入(仅支持大小不超过2M)',
  最大长度128个字符: '最大长度128个字符',
  '只允许包含中文、英文、数字、下划线 (_)、连字符 (-)，并且必须以中文、英文、数字开头和结尾': '只允许包含中文、英文、数字、下划线 (_)、连字符 (-)，并且必须以中文、英文、数字开头和结尾',
  配置项值不为数字: '配置项值不为数字',
  结果预览: '结果预览',
  已选择导入: '已选择导入',
  '个模板套餐，可按需要切换模板版本': '个模板套餐，可按需要切换模板版本',
  暂无预览: '暂无预览',
  请先从左侧选择导入的模板套餐: '请先从左侧选择导入的模板套餐',
  '检测到模板冲突，无法导入': '检测到模板冲突，无法导入',
  导入: '导入',
  配置文件导入成功: '配置文件导入成功',
  '检测到模板冲突，请先删除冲突套餐': '检测到模板冲突，请先删除冲突套餐',
  模板名称: '模板名称',
  模板路径: '模板路径',
  版本号: '版本号',
  该套餐下暂无模板: '该套餐下暂无模板',
  批量上传配置文件: '批量上传配置文件',
  上传配置文件包: '上传配置文件包',
  '支持扩展名：.zip  .tar  .gz': '支持扩展名：.zip  .tar  .gz',
  上传文件: '上传文件',
  '确认放弃下方修改，重新上传配置项包？': '确认放弃下方修改，重新上传配置项包？',
  重新上传: '重新上传',
  上传中: '上传中',
  共将导入: '共将导入',
  '个配置项，其中': '个配置项，其中',
  '个已存在,导入后将': '个已存在,导入后将',
  覆盖原配置: '覆盖原配置',
  去上传: '去上传',
  导入配置文件成功: '导入配置文件成功',
  已存在配置文件: '已存在配置文件',
  批量设置配置文件描述: '批量设置配置文件描述',
  批量设置文件权限: '批量设置文件权限',
  '只能输入三位 0~7 数字且文件own必须有读取权限': '只能输入三位 0~7 数字且文件own必须有读取权限',
  批量设置用户: '批量设置用户',
  批量设置用户组: '批量设置用户组',
  搜索模板套餐: '搜索模板套餐',
  管理模板: '管理模板',
  导入配置模板套餐时无法移除已有配置模板套餐: '导入配置模板套餐时无法移除已有配置模板套餐',
  '该套餐中没有可用配置文件，无法被导入到服务配置中': '该套餐中没有可用配置文件，无法被导入到服务配置中',
  配置文件内容: '配置文件内容',
  仅支持大小不超过: '仅支持大小不超过',
  分隔符: '分隔符',
  搜索: '搜索',
  全屏: '全屏',
  退出全屏: '退出全屏',
  '按 Esc 即可退出全屏模式': '按 Esc 即可退出全屏模式',
  请检查是否已正确使用分隔符: '请检查是否已正确使用分隔符',
  '类型必须为 string 或者 number': '类型必须为 string 或者 number',
  '类型为number 值不为number': '类型为number 值不为number',
  value不能为空: 'value不能为空',
  空字符: '空字符',
  自定义: '自定义',
  '您输入的分隔符错误,请重新输入': '您输入的分隔符错误,请重新输入',
  '您输入的分隔符过长,请重新输入': '您输入的分隔符过长,请重新输入',
  设置变量: '设置变量',
  设置变量成功: '设置变量成功',
  恢复默认值: '恢复默认值',
  '如果以下变量存在于全局变量中，其值将被重置为全局变量的默认值': '如果以下变量存在于全局变量中，其值将被重置为全局变量的默认值',
  暂无数据: '暂无数据',
  变量名称: '变量名称',
  类型: '类型',
  变量值: '变量值',
  变量说明: '变量说明',
  被引用: '被引用',
  不是数字类型: '不是数字类型',
  查看变量: '查看变量',
  关闭: '关闭',
  配置项值预览: '配置项值预览',
  修改人: '修改人',
  修改时间: '修改时间',
  变更状态: '变更状态',
  查看: '查看',
  对比: '对比',
  '确认删除该配置项？': '确认删除该配置项？',
  删除配置项成功: '删除配置项成功',
  恢复配置项成功: '恢复配置项成功',
  '一旦删除，该操作将无法撤销，请谨慎操作': '一旦删除，该操作将无法撤销，请谨慎操作',
  '配置项删除后，可以通过恢复按钮撤销删除': '配置项删除后，可以通过恢复按钮撤销删除',
  文本文件: '文本文件',
  二进制文件: '二进制文件',
  文本: '文本',
  二进制: '二进制',
  新增: '新增',
  修改: '修改',
  灰度中: '灰度中',
  搜索结果为空: '搜索结果为空',
  '可以尝试 调整关键词 或': '可以尝试 调整关键词 或',
  清空筛选条件: '清空筛选条件',
  编辑配置项: '编辑配置项',
  编辑配置项成功: '编辑配置项成功',
  查看配置项: '查看配置项',
  线上版本: '线上版本',
  对比版本: '对比版本',
  当前版本: '当前版本',
  上线版本: '上线版本',
  只查看差异文件: '只查看差异文件',
  搜索配置文件名称: '搜索配置文件名称',
  没有差异配置文件: '没有差异配置文件',
  只查看差异项: '只查看差异项',
  搜索配置项名称: '搜索配置项名称',
  没有差异配置项: '没有差异配置项',
  '前/后置脚本': '前/后置脚本',
  该版本下文件不存在: '该版本下文件不存在',
  文件已被删除: '文件已被删除',
  文件属性: '文件属性',
  文件内容: '文件内容',
  变化: '变化',
  配置模板版本: '配置模板版本',
  配置路径: '配置路径',
  移除套餐: '移除套餐',
  替换版本: '替换版本',
  '确认删除该配置文件？': '确认删除该配置文件？',
  '确认移除该配置模板套餐？': '确认移除该配置模板套餐？',
  配置模板套餐: '配置模板套餐',
  '移除后本服务配置将不再引用该配置模板套餐，以后需要时可以重新从配置模板导入': '移除后本服务配置将不再引用该配置模板套餐，以后需要时可以重新从配置模板导入',
  非模板配置: '非模板配置',
  移除模板套餐成功: '移除模板套餐成功',
  删除配置文件成功: '删除配置文件成功',
  编辑配置文件: '编辑配置文件',
  目标版本: '目标版本',
  当前最新为: '当前最新为',
  模板版本更新成功: '模板版本更新成功',
  查看配置文件: '查看配置文件',
  '配置项名/创建人/修改人': '配置项名/创建人/修改人',
  生成版本: '生成版本',
  版本生成中: '版本生成中',
  版本信息: '版本信息',
  同时上线版本: '同时上线版本',
  服务变量赋值: '服务变量赋值',
  '仅允许使用中文、英文、数字、下划线、中划线、点，且必须以中文、英文、数字开头和结尾': '仅允许使用中文、英文、数字、下划线、中划线、点，且必须以中文、英文、数字开头和结尾',
  编辑中: '编辑中',
  全部实例: '全部实例',
  默认分组: '默认分组',
  已上线实例: '已上线实例',
  除以下分组之外的所有实例: '除以下分组之外的所有实例',
  对比并上线: '对比并上线',
  版本已上线: '版本已上线',
  请选择上线分组: '请选择上线分组',
  本次上线分组: '本次上线分组',
  上线说明: '上线说明',
  确定上线: '确定上线',
  前置脚本: '前置脚本',
  后置脚本: '后置脚本',
  移除: '移除',
  上传: '上传',
  请至少选择一个排除分组实例: '请至少选择一个排除分组实例',
  选择上线实例范围: '选择上线实例范围',
  全部实例上线: '全部实例上线',
  排除分组实例上线: '排除分组实例上线',
  选择分组实例上线: '选择分组实例上线',
  '其它版本没有上线任何分组（默认版本除外），无法使用此选项': '其它版本没有上线任何分组（默认版本除外），无法使用此选项',
  全选: '全选',
  全不选: '全不选',
  暂无已上线的可选版本: '暂无已上线的可选版本',
  '搜索分组名称/标签key': '搜索分组名称/标签key',
  已上线分组不可取消选择: '已上线分组不可取消选择',
  上线预览: '上线预览',
  '上线后，以下分组将从以下各版本更新至当前版本': '上线后，以下分组将从以下各版本更新至当前版本',
  请先从左侧选择待上线的分组范围: '请先从左侧选择待上线的分组范围',
  首次上线: '首次上线',
  按版本选择: '按版本选择',
  共: '共',
  个分组: '个分组',
  变更版本: '变更版本',
  调整分组上线: '调整分组上线',
  该脚本未上线: '该脚本未上线',
  确定: '确定',
  预览: '预览',
  保存设置: '保存设置',
  脚本预览: '脚本预览',
  '<不使用脚本>': '<不使用脚本>',
  初始化脚本设置成功: '初始化脚本设置成功',
  确认恢复该版本: '确认恢复该版本',
  此操作会把改版本恢复至可用版本列表: '此操作会把改版本恢复至可用版本列表',
  版本恢复成功: '版本恢复成功',
  确认删除该版本: '确认删除该版本',
  版本删除成功: '版本删除成功',
  请输入关键字: '请输入关键字',
  '格式：': '格式:',
  'key 类型 value': 'key 类型 value',
  新增服务密钥: '新增服务密钥',
  调整分组上线成功: '调整分组上线成功',
  导入方式: '导入方式',
  文本导入: '文本导入',
  文件导入: '文件导入',
  '只支持string、number类型,其他类型请使用文件导入': '只支持string、number类型,其他类型请使用文件导入',
  '支持 JSON、YAML 等类型文件，后台会自动检测文件类型，配置项格式请参照': '支持 JSON、YAML 等类型文件，后台会自动检测文件类型，配置项格式请参照',
  示例文件包: '示例文件包',
  文件导入成功: '文件导入成功',
  文本导入成功: '文本导入成功',
  '解析失败，配置项格式不正确': '解析失败，配置项格式不正确',

  请选择: '请选择',
  配置项值已复制: '配置项值已复制',
  // 分组管理
  新增分组: '新增分组',
  按标签分类查看: '按标签分类查看',
  '分组名称/标签选择器': '分组名称/标签选择器',
  分组名称: '分组名称',
  标签选择器: '标签选择器',
  服务可见范围: '服务可见范围',
  公开: '公开',
  分组状态: '分组状态',
  上线服务数: '上线服务数',
  编辑分组: '编辑分组',
  '确认删除该分组？': '确认删除该分组？',
  '分组由 1 个或多个标签选择器组成，服务配置版本选择分组上线结合客户端配置的标签用于灰度发布、A/B Test等运营场景，详情参考文档：': '分组由 1 个或多个标签选择器组成，服务配置版本选择分组上线结合客户端配置的标签用于灰度发布、A/B Test等运营场景，详情参考文档：',
  删除分组成功: '删除分组成功',
  '分组已上线，不能': '分组已上线，不能',
  创建分组成功: '创建分组成功',
  请输入分组名称: '请输入分组名称',
  指定服务: '指定服务',
  请选择服务: '请选择服务',
  '标签选择器由key、操作符、value组成，筛选符合条件的客户端拉取服务配置，一般用于灰度发布服务配置': '标签选择器由key、操作符、value组成，筛选符合条件的客户端拉取服务配置，一般用于灰度发布服务配置',
  分组规则表单不能为空: '分组规则表单不能为空',
  '仅允许使用中文、英文、数字、下划线、中划线，且必须以中文、英文、数字开头和结尾': '仅允许使用中文、英文、数字、下划线、中划线，且必须以中文、英文、数字开头和结尾',
  指定服务不能为空: '指定服务不能为空',
  编辑分组成功: '编辑分组成功',
  上线服务: '上线服务',
  '服务名称/服务版本': '服务名称/服务版本',
  服务版本: '服务版本',

  // 全局变量
  配置模板与变量: '配置模板与变量',
  新增变量: '新增变量',
  导入变量: '导入变量',
  请输入变量名称: '请输入变量名称',
  默认值: '默认值',
  描述: '描述',
  ' 确认删除该全局变量？': '确认删除该全局变量？',
  '一旦删除，该操作将无法撤销，服务配置文件中不可再引用该全局变量，请谨慎操作': '一旦删除，该操作将无法撤销，服务配置文件中不可再引用该全局变量，请谨慎操作',
  '定义全局变量后可供业务下所有的服务配置文件引用，使用go template语法引用，例如{{ .bk_bscp_appid }},变量使用详情请参考：': '定义全局变量后可供业务下所有的服务配置文件引用，使用go template语法引用，例如{{ .bk_bscp_appid }},变量使用详情请参考：',
  删除变量成功: '删除变量成功',
  创建: '创建',
  创建变量成功: '创建变量成功',
  变量名称不能为空: '变量名称不能为空',
  '仅允许使用中文、英文、数字、下划线，且不能以数字开头': '仅允许使用中文、英文、数字、下划线，且不能以数字开头',
  '无效默认值，类型为number值不为数字': '无效默认值，类型为number值不为数字',
  编辑变量: '编辑变量',
  编辑变量成功: '编辑变量成功',
  变量内容: '变量内容',
  变量必须以bk_bscp_或BK_BSCP_开头: '变量必须以bk_bscp_或BK_BSCP_开头',
  导入变量成功: '导入变量成功',
  '示例：': '示例：',
  '变量名 变量类型 变量值 变量描述（可选）': '变量名 变量类型 变量值 变量描述（可选）',
  ' bk_bscp_nginx_port number 8080 nginx端口': ' bk_bscp_nginx_port number 8080 nginx端口',

  // 配置模板
  '配置模板用于统一业务下服务间配置文件复用，可以在创建服务配置时引用配置模板。': '配置模板用于统一业务下服务间配置文件复用，可以在创建服务配置时引用配置模板。',
  搜索空间: '搜索空间',
  创建空间: '创建空间',
  '确认删除该配置模板空间？': '确认删除该配置模板空间？',
  '空间可将业务下不同使用场景的配置模板文件隔离，每个空间内的配置文件路径+配置文件名称是唯一的，每个业务下会自动创建一个默认空间': '空间可将业务下不同使用场景的配置模板文件隔离，每个空间内的配置文件路径+配置文件名称是唯一的，每个业务下会自动创建一个默认空间',
  未能删除: '未能删除',
  请先确认删除此空间下所有配置文件: '请先确认删除此空间下所有配置文件',
  请先确认删除此空间下所有配置套餐: '请先确认删除此空间下所有配置套餐',
  删除空间成功: '删除空间成功',
  新建空间: '新建空间',
  模板空间名称: '模板空间名称',
  模板空间描述: '模板空间描述',
  创建空间成功: '创建空间成功',
  编辑空间: '编辑空间',
  编辑空间成功: '编辑空间成功',
  新建模板套餐: '新建模板套餐',
  全部配置文件: '全部配置文件',
  未指定套餐: '未指定套餐',
  克隆: '克隆',
  克隆模板套餐: '克隆模板套餐',
  克隆成功: '克隆成功',
  创建成功: '创建成功',
  '确认删除该配置模板套餐？': '确认删除该配置模板套餐？',
  '一旦删除，该操作将无法撤销，以下服务配置的未命名版本中引用该套餐的内容也将清除': '一旦删除，该操作将无法撤销，以下服务配置的未命名版本中引用该套餐的内容也将清除',
  暂无未命名版本引用此套餐: '暂无未命名版本引用此套餐',
  引用此套餐的服务: '引用此套餐的服务',
  删除配置模板套餐成功: '删除配置模板套餐成功',
  编辑模板套餐: '编辑模板套餐',
  编辑成功: '编辑成功',
  模板套餐名称: '模板套餐名称',
  模板套餐描述: '模板套餐描述',
  '提醒：修改可见范围后，服务': '提醒：修改可见范围后，服务',
  将不再引用此套餐: '将不再引用此套餐',
  使用模板: '使用模板',
  新服务中使用: '新服务中使用',
  当前使用此套餐的服务: '当前使用此套餐的服务',
  '服务名称/配置文件版本': '服务名称/配置文件版本',
  配置文件版本: '配置文件版本',
  引用此配置文件的服务: '引用此配置文件的服务',
  '配置文件名称/路径/描述/创建人/更新人': '配置文件名称/路径/描述/创建人/更新人',
  所在套餐: '所在套餐',
  更新人: '更新人',
  更新时间: '更新时间',
  版本管理: '版本管理',
  添加至套餐: '添加至套餐',
  移出套餐: '移出套餐',
  添加至: '添加至',
  添加配置文件: '添加配置文件',
  添加已有配置文件: '添加已有配置文件',
  去创建: '去创建',
  新建配置文件成功: '新建配置文件成功',
  创建至套餐: '创建至套餐',
  模板套餐: '模板套餐',
  使用此套餐的服务: '使用此套餐的服务',
  '若未指定套餐，此配置文件模板将无法被服务引用。后续请使用「添加至」或「添加已有配置文件」功能添加至指定套餐': '若未指定套餐，此配置文件模板将无法被服务引用。后续请使用「添加至」或「添加已有配置文件」功能添加至指定套餐',
  以下服务配置的未命名版本引用目标套餐的内容也将更新: '以下服务配置的未命名版本引用目标套餐的内容也将更新',
  从已有配置文件添加: '从已有配置文件添加',
  '配置文件名称/路径/描述': '配置文件名称/路径/描述',
  已选: '已选',
  请先从左侧选择配置文件: '请先从左侧选择配置文件',
  添加: '添加',
  添加配置文件成功: '添加配置文件成功',
  上传至套餐: '上传至套餐',
  批量添加至: '批量添加至',
  添加至模板套餐: '添加至模板套餐',
  以下服务配置的未命名版本中将添加已选配置文件的: '以下服务配置的未命名版本中将添加已选配置文件的',
  目标模板套餐: '目标模板套餐',
  批量删除: '批量删除',
  '确认删除以下配置文件？': '确认删除以下配置文件？',
  批量移出: '批量移出',
  批量移出当前套餐: '批量移出当前套餐',
  确定移出: '确定移出',
  所在模板套餐: '所在模板套餐',
  以下服务配置的未命名版本中引用此套餐的内容也将更新: '以下服务配置的未命名版本中引用此套餐的内容也将更新',
  配置文件移出套餐成功: '配置文件移出套餐成功',
  '确认从配置套餐中移出该配置文件?': '确认从配置套餐中移出该配置文件?',
  当前: '当前',
  '移出后配置文件将不存在任一套餐。你仍可在「全部配置文件」或「未指定套餐」分类下找回。': '移出后配置文件将不存在任一套餐。你仍可在「全部配置文件」或「未指定套餐」分类下找回。',
  确认移出: '确认移出',
  移出套餐成功: '移出套餐成功',
  新建版本: '新建版本',
  '版本号/版本说明/更新人': '版本号/版本说明/更新人',
  选择载入版本: '选择载入版本',
  版本说明: '版本说明',
  展开列表: '展开列表',
  '支持扩展名：.bin，文件大小100M以内': '支持扩展名：.bin，文件大小100M以内',
  创建版本成功: '创建版本成功',
  '确认更新配置文件版本？': '确认更新配置文件版本？',
  以下套餐及服务未命名版本中引用的此配置文件也将更新: '以下套餐及服务未命名版本中引用的此配置文件也将更新',
  引用此模板的服务: '引用此模板的服务',

  // 脚本管理
  全部脚本: '全部脚本',
  未分类: '未分类',
  新建脚本: '新建脚本',
  脚本名称: '脚本名称',
  脚本语言: '脚本语言',
  分类标签: '分类标签',
  '确认删除该脚本？': '确认删除该脚本？',
  脚本: '脚本',
  '一旦删除，该操作将无法撤销，以下服务配置的未命名版本中引用该脚本也将清除': '一旦删除，该操作将无法撤销，以下服务配置的未命名版本中引用该脚本也将清除',
  暂无未命名版本引用此脚本: '暂无未命名版本引用此脚本',
  引用此脚本的服务: '引用此脚本的服务',
  删除版本成功: '删除版本成功',
  请选择标签或输入新标签按Enter结束: '请选择标签或输入新标签按Enter结束',
  脚本描述: '脚本描述',
  脚本内容: '脚本内容',
  不能超过64个字符: '不能超过64个字符',
  脚本创建成功: '脚本创建成功',
  关联配置文件: '关联配置文件',
  '服务名称/版本名称/被引用的版本': '服务名称/版本名称/被引用的版本',
  脚本版本: '脚本版本',
  复制并新建: '复制并新建',
  '确认删除该脚本版本？': '确认删除该脚本版本？',
  '确定上线此版本？': '确定上线此版本？',
  '上线后，之前的线上版本将被置为「已下线」状态': '上线后，之前的线上版本将被置为「已下线」状态',
  '当前已有「未上线」版本': '当前已有「未上线」版本',
  前往编辑: '前往编辑',
  创建版本: '创建版本',
  选择载入脚本: '选择载入脚本',
  '无效名称，只允许包含中文、英文、数字、下划线()、连字符(-)、空格，且必须以中文、英文、数字开头和结尾': '无效名称，只允许包含中文、英文、数字、下划线()、连字符(-)、空格，且必须以中文、英文、数字开头和结尾',
  编辑版本: '编辑版本',
  脚本内容不能为空: '脚本内容不能为空',
  编辑版本成功: '编辑版本成功',
  新建版本成功: '新建版本成功',
  上线: '上线',
  已下线: '已下线',
  脚本类型: '脚本类型',

  // 服务密钥
  '密钥仅用于 SDK/API 拉取配置使用。服务管理/配置管理/分组管理等功能的权限申请，请前往': '密钥仅用于 SDK/API 拉取配置使用。服务管理/配置管理/分组管理等功能的权限申请，请前往',
  蓝鲸权限中心: '蓝鲸权限中心',
  新建密钥: '新建密钥',
  '密钥名称/说明/更新人': '密钥名称/说明/更新人',
  密钥名称: '密钥名称',
  密钥名称支持中英文: '密钥名称支持中英文',
  密钥: '密钥',
  待确认: '待确认',
  说明: '说明',
  请输入密钥说明: '请输入密钥说明',
  关联规则: '关联规则',
  已启用: '已启用',
  已禁用: '已禁用',
  关联服务配置: '关联服务配置',
  '确认删除密钥？': '确认删除密钥？',
  删除的密钥: '删除的密钥',
  无法找回: '无法找回',
  ',请谨慎操作！': ',请谨慎操作！',
  请输入密钥名称: '请输入密钥名称',
  以确认删除: '以确认删除',
  服务密钥已复制: '服务密钥已复制',
  新建服务密钥成功: '新建服务密钥成功',
  密钥名称修改成功: '密钥名称修改成功',
  密钥说明修改成功: '密钥说明修改成功',
  确定禁用此密钥: '确定禁用此密钥',
  '禁用密钥后，使用此密钥的应用将无法正常使用 SDK/API 拉取配置': '禁用密钥后，使用此密钥的应用将无法正常使用 SDK/API 拉取配置',
  禁用: '禁用',
  禁用成功: '禁用成功',
  启用成功: '启用成功',
  删除服务密钥成功: '删除服务密钥成功',
  '已启用，不能删除': '已启用，不能删除',
  无效名称: '无效名称',
  '只允许包含中文、英文、数字、下划线 (_)、连字符 (-)，并且必须以中文、英文、数字开头和结尾。': '只允许包含中文、英文、数字、下划线 (_)、连字符 (-)，并且必须以中文、英文、数字开头和结尾。',
  编辑规则: '编辑规则',
  编辑规则成功: '编辑规则成功',
  '文件型配置，以选择服务myservice为例:': '文件型配置，以选择服务myservice为例:',
  '键值（KV）型配置，以选择服务myservice为例:': '键值（KV）型配置，以选择服务myservice为例:',
  '关联myservice服务下所有的配置(包含子目录)': '关联myservice服务下所有的配置(包含子目录)',
  '关联myservice服务/etc目录下所有的配置(不含子目录)': '关联myservice服务/etc目录下所有的配置(不含子目录)',
  '关联myservice服务/etc/nginx/nginx.conf文件': '关联myservice服务/etc/nginx/nginx.conf文件',
  关联myservice服务下所有配置项: '关联myservice服务下所有配置项',
  关联myservice服务下所有以demo_开头的配置项: '关联myservice服务下所有以demo_开头的配置项',
  共有: '共有',
  项关联规则: '项关联规则',
  暂未设置关联规则: '暂未设置关联规则',
  配置关联规则: '配置关联规则',
  查看规则示例: '查看规则示例',
  撤销本次删除: '撤销本次删除',
  '输入的规则有误，请重新确认': '输入的规则有误，请重新确认',
  请输入文件路径: '请输入文件路径',
  请输入配置项名称: '请输入配置项名称',

  // 公共组件
  页面不存在: '页面不存在',
  产品仅供内部体验: '产品仅供内部体验',
  '如需体验，请联系': '如需体验，请联系',
  '你没有相应业务的访问权限，请前往申请相关业务权限': '你没有相应业务的访问权限，请前往申请相关业务权限',
  无: '无',
  权限: '权限',
  无访问权限: '无访问权限',
  确认彻底删除: '确认彻底删除',
  条配置文件: '条配置文件',
  '删除后不可找回，请谨慎操作。': '删除后不可找回，请谨慎操作。',
  确认删除: '确认删除',
  '确认离开当前页？': '确认离开当前页？',
  离开会导致未保存信息丢失: '离开会导致未保存信息丢失',
  离开: '离开',
  需要申请的权限: '需要申请的权限',
  关联的资源实例: '关联的资源实例',
  无数据: '无数据',
  已申请: '已申请',
  去申请: '去申请',
  产品功能特性: '产品功能特性',
  技术支持: '技术支持',
  社区论坛: '社区论坛',
  产品官网: '产品官网',
  请输入: '请输入',
  确认: '确认',
};
