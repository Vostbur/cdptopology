(function (nx) {
    /**
     * Приложение на NeXt UI
     */
    // Инициализация топологии
    var topo = new nx.graphic.Topology({
        // Ширина и высота view приложения
        width: 1200,
        height: 700,
        // Процессор данных, отвечает за расстановку нод.
        // 'force' стремится расставить ноды на равном 
        // удалении друг от друга. 'quick' расставляет их
        // в произвольных местах
        dataProcessor: 'force',
        // уникальный идентификатор нод и линков
        identityKey: 'id',
        // Конфигурация нод
        nodeConfig: {
            label: 'model.name',
            iconType:'model.icon',
        },
        // Конфигурация линков
        linkConfig: {
            // Отображение множественных линков дугами,
            // можно поменять на 'parallel'
            linkType: 'curve',
        },
        // Отображать пиктограммы нод, при false отрисует точку
        showIcon: true,
    });

    var Shell = nx.define(nx.ui.Application, {
        methods: {
            start: function () {
                // записать данные топологии из переменной
                topo.data(topologyData);
                // и прикрепить их к документу
                topo.attach(this);
            }
        }
    });

    // создать инстанс приложения
    var shell = new Shell();
    // запустить приложение
    shell.start();
})(nx);
