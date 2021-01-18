package instruction

import (
	"errors"

	"github.com/bitwormhole/go-wormhole-core/collection"

	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

type defaultGitInitRunner struct {
	task *GitInit
}

func (inst *defaultGitInitRunner) Run() error {

	bare := inst.task.Bare
	path := inst.task.Path
	name := inst.task.Name

	dir1 := path.GetChild(name)
	var repoDir fs.Path

	if bare {
		repoDir = dir1
	} else {
		repoDir = dir1.GetChild(".git")
	}

	if dir1.Exists() {
		return errors.New("the target directory is exists: " + dir1.Path())
	}

	if repoDir.Exists() {
		return errors.New("the repository directory is exists: " + repoDir.Path())
	}

	err := repoDir.Mkdirs()
	if err != nil {
		return err
	}

	err = inst.mkfiles(repoDir)
	if err != nil {
		return err
	}

	err = inst.initConfig(repoDir)
	if err != nil {
		return err
	}

	err = inst.initDescription(repoDir)
	if err != nil {
		return err
	}

	err = inst.initHEAD(repoDir)
	if err != nil {
		return err
	}

	return nil
}

func (inst *defaultGitInitRunner) mkfiles(repo fs.Path) error {

	files := []string{}
	dirs := []string{}

	dirs = append(dirs, "objects/pack")
	dirs = append(dirs, "objects/info")
	dirs = append(dirs, "refs/heads")
	dirs = append(dirs, "refs/tags")
	dirs = append(dirs, "logs")
	dirs = append(dirs, "info")
	dirs = append(dirs, "hooks")

	files = append(files, "config")
	files = append(files, "HEAD")
	files = append(files, "description")

	// make dirs
	for index := range dirs {
		name := dirs[index]
		dir := repo.GetChild(name)
		err := dir.Mkdirs()
		if err != nil {
			return err
		}
	}

	// make files
	for index := range files {
		name := files[index]
		dir := repo.GetChild(name)
		err := dir.CreateFile(nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *defaultGitInitRunner) initDescription(repo fs.Path) error {

	file := repo.GetChild("description")

	if !file.IsFile() {
		return errors.New("no file " + file.Path())
	}

	if file.Size() != 0 {
		return errors.New("target is not a empty file: " + file.Path())
	}

	text := "Repository: " + inst.task.Name
	return file.GetIO().WriteText(text, nil)
}

func (inst *defaultGitInitRunner) initHEAD(repo fs.Path) error {

	file := repo.GetChild("HEAD")

	if !file.IsFile() {
		return errors.New("no file " + file.Path())
	}

	if file.Size() != 0 {
		return errors.New("target is not a empty file: " + file.Path())
	}

	text := "ref: refs/heads/main\n"
	return file.GetIO().WriteText(text, nil)
}

func (inst *defaultGitInitRunner) initConfig(repo fs.Path) error {

	configfile := repo.GetChild("config")

	if !configfile.IsFile() {
		return errors.New("no file " + configfile.Path())
	}

	if configfile.Size() != 0 {
		return errors.New("target is not a empty file: " + configfile.Path())
	}

	const yes string = "true"
	const no string = "false"

	repositoryformatversion := "0"
	filemode := no
	bare := no
	logallrefupdates := yes
	symlinks := no
	ignorecase := yes

	if inst.task.Bare {
		bare = yes
	}

	config := &collection.SimpleProperties{}

	config.SetProperty("core.repositoryformatversion", repositoryformatversion)
	config.SetProperty("core.filemode", filemode)
	config.SetProperty("core.bare", bare)
	config.SetProperty("core.logallrefupdates", logallrefupdates)
	config.SetProperty("core.symlinks", symlinks)
	config.SetProperty("core.ignorecase", ignorecase)

	text := collection.FormatPropertiesWithSegment(config)
	fsio := configfile.GetIO()
	return fsio.WriteText(text, nil)
}
