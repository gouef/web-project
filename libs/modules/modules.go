package modules

import (
	"fmt"
	"github.com/gouef/finder"
	"github.com/gouef/utils"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var ModuleName string
var Controllers = map[string]Controller{}

var (
	namespaceCache   = make(map[uintptr]string)
	pathFileCache    = make(map[uintptr]string)
	ControllersCache = make(map[string]Controller)
	cacheMutex       sync.RWMutex
)

func GetName() string {
	if ModuleName != "" {
		return ModuleName
	}

	pc, _, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return ""
	}

	fullName := fn.Name()
	parts := strings.Split(fullName, "/")
	if len(parts) < 3 {
		return ""
	}

	ModuleName = strings.Join(parts[:3], "/")

	return ModuleName
}

func GetPathDir() string {
	return strings.Replace(GetPathFile(), ".go", "", 1)
}

func GetController() Controller {
	name := GetNamespace()
	cacheMutex.RLock()
	if cachedController, ok := ControllersCache[name]; ok {
		cacheMutex.RUnlock()
		fmt.Println("Using cached controller for:", name)
		return cachedController
	}
	cacheMutex.RUnlock()

	fullPath := GetPathFile()
	fileName := filepath.Base(fullPath)
	controllerName := strings.TrimSuffix(fileName, "Controller.go")
	namespace := strings.TrimPrefix(strings.TrimSuffix(fullPath, "/"+controllerName+"Controller.go"), "/")
	controller := Controller{
		Name:      controllerName,
		Namespace: namespace,
		Path:      fullPath,
		Method:    name,
	}

	cacheMutex.Lock()
	ControllersCache[name] = controller
	cacheMutex.Unlock()

	return Controllers[name]
}

func GetPathFile() string {
	var caller string
	modulePrefix := GetName()

	for i := 0; ; i++ {
		pc, file, _, ok := runtime.Caller(i)

		if !ok {
			break
		}

		cacheMutex.RLock()
		if cached, ok := pathFileCache[pc]; ok {
			cacheMutex.RUnlock()
			if cached != "" {
				caller = cached
			}
			continue
		}
		cacheMutex.RUnlock()

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		fullName := fn.Name()
		if strings.HasPrefix(fullName, modulePrefix) {
			caller = file
			cacheMutex.Lock()
			pathFileCache[pc] = fullName
			cacheMutex.Unlock()
		} else {
			cacheMutex.Lock()
			pathFileCache[pc] = ""
			cacheMutex.Unlock()
		}
	}

	return strings.TrimPrefix(caller, finder.GetProjectRoot())
}

func GetNamespace() string {
	var caller string
	modulePrefix := GetName()

	for i := 1; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}

		cacheMutex.RLock()
		if cached, ok := namespaceCache[pc]; ok {
			cacheMutex.RUnlock()
			if cached != "" {
				caller = cached
			}
			continue
		}
		cacheMutex.RUnlock()
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		fullName := fn.Name()
		if strings.HasPrefix(fullName, modulePrefix) {
			caller = fullName
			cacheMutex.Lock()
			namespaceCache[pc] = fullName
			cacheMutex.Unlock()
		} else {
			cacheMutex.Lock()
			namespaceCache[pc] = ""
			cacheMutex.Unlock()
		}
	}

	namespace := strings.TrimPrefix(caller, utils.Implode("", []string{modulePrefix, "/"}))

	return namespace
}
