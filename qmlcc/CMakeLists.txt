
cmake_minimum_required(VERSION 3.16)

project(QtQuickInline VERSION 1.0.1 LANGUAGES CXX C)


set(CMAKE_C_STANDARD 99)
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_FLAGS "-g -ggdb -O0 -DQT_QML_DEBUG -DQT_DECLARATIVE_DEBUG")

set(CMAKE_VERBOSE_MAKEFILE true)
set(CMAKE_CXX_EXTENSIONS off)


find_package(Qt6 REQUIRED COMPONENTS Core)
# find_package(Qt6 REQUIRED COMPONENTS Widgets)
# find_package(Qt6 REQUIRED COMPONENTS Svg)
# find_package(Qt6 REQUIRED COMPONENTS Xml)
# declaretive
find_package(Qt6 REQUIRED COMPONENTS Qml)
find_package(Qt6 REQUIRED COMPONENTS QmlIntegration)
find_package(Qt6 REQUIRED COMPONENTS Quick)
find_package(Qt6 REQUIRED COMPONENTS QuickControls2)
find_package(Qt6 REQUIRED COMPONENTS QuickTemplates2)
# find_package(Qt6 REQUIRED COMPONENTS QuickWidgets)
qt_standard_project_setup()

# find_library(theffi libffi ffi REQUIRED)
# message(${theffi})

set(MAINSRC
    qtuninline.cpp
)
qt_add_library(QtQuickInline SHARED ${MAINSRC})
target_link_libraries(QtQuickInline PRIVATE
    # Qt6::Widgets
    Qt6::Qml Qt6::QmlIntegration Qt6::Quick
    Qt6::QuickControls2 Qt6::QuickTemplates2
    # Qt6::QuickWidgets
    # 以下添加 private header
    Qt6::QuickPrivate Qt6::QuickControls2Private Qt6::QuickTemplates2Private
    # theffi
    # ${GOLIB}
)
