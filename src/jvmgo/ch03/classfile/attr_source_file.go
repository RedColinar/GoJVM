package classfile

type SourceFileAttribute struct {
	cp				ConstantPool
	//常量池索引，指向ConstantUtf8Info
	sourceFileIndex	uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader){
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string{
	return self.cp.getUtf8(self.sourceFileIndex )
}