package service

//var SimpleUploader = new(uploader)
//
//type uploader struct {
//	err       error
//	file      *os.File
//	fileInfos []os.FileInfo
//
//	chunkDir  string
//	checkPath string
//	chunkPath string
//	finishDir string
//}
//
////@author: [SliverHorn](https://github.com/SliverHorn)
////@description: 上传保存切片文件
//func (u *uploader) Upload(context *gin.Context, header *multipart.FileHeader, info *request.Upload) error {
//	u.chunkDir = global.Config.Uploader.GetIdentifier(info.Identifier)
//	if ok, _ := utils.PathExists(u.chunkDir); !ok {
//		if u.err = utils.CreateDir(u.chunkDir); u.err != nil {
//			zap.L().Error("创建目录失败!", zap.Any("err", u.err))
//		}
//	}
//	u.chunkPath = u.chunkDir + info.Filename + info.ChunkNumber
//	if u.err = context.SaveUploadedFile(header, u.chunkPath); u.err != nil {
//		return errors.New("切片创建失败! ")
//	}
//	return global.Db.Scopes(info.Create(u.chunkPath)).Error
//}
//
////@author: [SliverHorn](https://github.com/SliverHorn)
////@description: 保存文件切片路径
//func (u *uploader) CreateChunk(uploader model.SimpleUploader) error {
//	return global.Db.Create(uploader).Error
//}
//
////@author: [SliverHorn](https://github.com/SliverHorn)
////@description: 检查文件是否已经上传过
//func (u *uploader) CheckFileMd5(info *request.CheckFileMd5) (uploads *[]model.SimpleUploader, isDone bool, err error) {
//	var entity []model.SimpleUploader
//	err = global.Db.Find(&entity, "identifier = ? AND is_done = ?", info.Md5, false).Error
//	isDone = errors.Is(global.Db.First(&model.SimpleUploader{}, "identifier = ? AND is_done = ?", info.Md5, true).Error, gorm.ErrRecordNotFound)
//	return &entity, !isDone, err
//}
//
////@author: [SliverHorn](https://github.com/SliverHorn)
////@description: 合并文件
//func (u *uploader) MergeFileMd5(info *request.MergeFileMd5) error {
//	u.finishDir = global.Config.Uploader.FinishDir
//	u.checkPath = global.Config.Uploader.GetCheckPath(info.Md5)
//	if !errors.Is(global.Db.First(&model.SimpleUploader{}, "identifier = ? AND is_done = ?", info.Md5, true).Error, gorm.ErrRecordNotFound) { //如果文件上传成功 不做后续操作 通知成功即可
//		return nil
//	}
//
//	if u.fileInfos, u.err = ioutil.ReadDir(u.checkPath); u.err != nil { // 打开切片文件夹
//		return u.err
//	}
//	_ = os.MkdirAll(u.finishDir, os.ModePerm)
//
//	if u.file, u.err = os.OpenFile(u.finishDir+info.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); u.err != nil { // 创建目标文件
//		return u.err
//	}
//
//	defer func() { //关闭文件
//		_ = u.file.Close()
//	}()
//
//	for k := range u.fileInfos { // 将切片文件按照顺序写入
//		content, _ := ioutil.ReadFile(u.checkPath + "/" + info.Filename + strconv.Itoa(k+1))
//		if _, u.err = u.file.Write(content); u.err != nil {
//			_ = os.Remove(u.finishDir + info.Filename)
//		}
//	}
//	if u.err = global.Db.Transaction(func(tx *gorm.DB) error {
//		if u.err = tx.Delete(&model.SimpleUploader{}, "identifier = ? AND is_done = ?", info.Md5, false).Error; u.err != nil { // 删除切片信息
//			return u.err
//		}
//		var entity = model.SimpleUploader{IsDone: true, FilePath: u.finishDir + info.Filename, Filename: info.Filename, Identifier: info.Md5}
//		if u.err = tx.Create(&entity).Error; u.err != nil { // 添加文件信息
//			return u.err
//		}
//		return nil
//	}); u.err != nil {
//		return u.err
//	}
//	return os.RemoveAll(u.checkPath) //清除切片
//}
