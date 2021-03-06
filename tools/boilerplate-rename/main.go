package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var steps int = 2
var oldPackage string = "co.touchlab.kampkit"
var newPackage string = "be.afelio.boilerplate"

// Mandatory to work
var android bool = false
var ios bool = false

func main() {
	parseArgs()

	if android {
		replacePackageOnFiles(getFilesToUpdate(), oldPackage, newPackage)
		moveFilesToNewPackage(getPackagesPaths(), strings.ReplaceAll(oldPackage, ".", "/"), strings.ReplaceAll(newPackage, ".", "/"))
		removeOldPackages(getPackagesPathsToDelete())
	}

	if ios {
		log.Println("iOS stuffs")
	}
}

func parseArgs() {
	args := os.Args[1:]

	for i, arg := range args {
		if arg == "--android" {
			android = true
		}
		if arg == "--ios" {
			ios = true
		}
		if arg == "-o" && i+1 < len(args) {
			oldPackage = args[i+1]
		}
		if arg == "-n" && i+1 < len(args) {
			newPackage = args[i+1]
		}
		if arg == "--step" && i+1 < len(args) {
			s, err := strconv.ParseInt(args[i+1], 10, 32)
			if err != nil {
				panic(err)
			}
			steps = int(s)
		}
	}
}

func up(step int) string {
	steps := ""
	for i := 0; i < step; i++ {
		steps = steps + "../"
	}
	return steps
}

func replacePackageOnFiles(paths []string, oldPackage string, newPackage string) {

	for _, path := range paths {
		read, err := ioutil.ReadFile(up(steps) + path)
		if err != nil {
			log.Println(err)
			return
		}

		newContents := strings.Replace(string(read), oldPackage, newPackage, -1)

		err = ioutil.WriteFile(up(steps)+path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Old package", oldPackage, "replaced by", newPackage, "in", len(paths), "files")
	fmt.Println("")
}

func updatePath(currentPath string, oldPackage string, newPackage string) string {
	return strings.Replace(currentPath, oldPackage, newPackage, 1)
}

// Method to recreate new package and move previous existing file into it
func moveFilesToNewPackage(paths []string, oldPath string, newPath string) {
	log.Println("Task: Create new package structure and move files")
	log.Println()

	filesMoved := 0

	for _, path := range paths {
		updatedPath := updatePath(path, oldPath, newPath)
		err := os.MkdirAll(up(steps)+updatedPath, 0755)

		if err != nil {
			panic(err)
		}

		errr := filepath.Walk(up(steps)+path,
			func(subpath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				// Ignore top level path
				if subpath == up(steps)+path {
					return nil
				}

				// Change path
				newPath := strings.Replace(subpath, oldPath, newPath, 1)

				if info.IsDir() {
					err := os.MkdirAll(newPath, 0755)
					if err != nil {
						panic(err)
					}
					log.Println("Directory", info.Name(), "created to", newPath)
				} else {
					read, err := ioutil.ReadFile(subpath)
					if err != nil {
						panic(err)
					}

					err = ioutil.WriteFile(newPath, []byte(read), 0755)
					if err != nil {
						panic(err)
					}

					filesMoved += 1
					log.Println("File", info.Name(), "moved to", newPath)
				}

				return nil
			})
		if errr != nil {
			log.Println(err)
		}
	}
	log.Println()
	log.Println("SUCCESS:", filesMoved, "files moved")
}

func removeOldPackages(paths []string) {
	log.Println("Task: Remove old packages")
	for _, path := range paths {
		err := os.RemoveAll(up(steps) + path)

		if err != nil {
			panic(err)
		}
	}

	log.Println("SUCCESS:", len(paths), "old packages deleted")
	log.Println()
}

func getFilesToUpdate() []string {
	return []string{
		"app/build.gradle.kts",
		"app/src/main/AndroidManifest.xml",
		"app/src/main/java/co/touchlab/kampkit/android/BreedViewModel.kt",
		"app/src/main/java/co/touchlab/kampkit/android/MainActivity.kt",
		"app/src/main/java/co/touchlab/kampkit/android/MainApp.kt",
		"app/src/main/java/co/touchlab/kampkit/android/adapter/MainAdapter.kt",
		"app/src/main/java/co/touchlab/kampkit/android/adapter/MainViewHolder.kt",
		"app/src/main/res/layout/activity_main.xml",
		"ios/KaMPKitiOS.xcodeproj/project.pbxproj",
		"shared/build.gradle.kts",
		"shared/src/androidMain/kotlin/co/touchlab/kampkit/KoinAndroid.kt",
		"shared/src/androidMain/kotlin/co/touchlab/kampkit/PlatformAndroid.kt",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit/BaseTest.kt",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit/CoroutineTestRule.kt",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit/KoinTest.kt",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit/TestUtilAndroid.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/AppInfo.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/DatabaseHelper.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/Koin.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/Platform.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/ktor/DogApiImpl.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/ktor/KtorApi.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/models/BreedModel.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/models/DataState.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/response/BreedResult.kt",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit/sqldelight/CoroutinesExtensions.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/BaseTest.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/BreedModelTest.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/ConcurrencyTest.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/SqlDelightTest.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/TestAppInfo.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/TestUtil.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/mock/ClockMock.kt",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit/mock/KtorApiMock.kt",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit/KoinIOS.kt",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit/MainScope.kt",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit/NativeViewModel.kt",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit/PlatformiOS.kt",
		"shared/src/iosTest/kotlin/co/touchlab/kampkit/BaseTest.kt",
		"shared/src/iosTest/kotlin/co/touchlab/kampkit/KoinTest.kt",
		"shared/src/iosTest/kotlin/co/touchlab/kampkit/TestUtilIOS.kt",
		"shared/src/main/AndroidManifest.xml",
	}
}

func getPackagesPaths() []string {
	return []string{
		"app/src/main/java/co/touchlab/kampkit",
		"shared/src/androidMain/kotlin/co/touchlab/kampkit",
		"shared/src/androidTest/kotlin/co/touchlab/kampkit",
		"shared/src/commonMain/kotlin/co/touchlab/kampkit",
		"shared/src/commonMain/sqldelight/co/touchlab/kampkit",
		"shared/src/commonTest/kotlin/co/touchlab/kampkit",
		"shared/src/iosMain/kotlin/co/touchlab/kampkit",
		"shared/src/iosTest/kotlin/co/touchlab/kampkit",
	}
}

func getPackagesPathsToDelete() []string {
	return []string{
		"app/src/main/java/co",
		"shared/src/androidMain/kotlin/co",
		"shared/src/androidTest/kotlin/co",
		"shared/src/commonMain/kotlin/co",
		"shared/src/commonMain/sqldelight/co",
		"shared/src/commonTest/kotlin/co",
		"shared/src/iosMain/kotlin/co",
		"shared/src/iosTest/kotlin/co",
	}
}
