<script setup>
import { defineProps, defineEmits } from 'vue';
import { useFilterStore } from '@/stores/filterStore';

const props = defineProps({
    active: {
        type: Boolean,
    },
    activeObject: {
        type: String,
    }
})

const emit = defineEmits(['closeFilter', 'getMarksWithFilter'])

const closeFilter = () => {
    emit('closeFilter')
}

const getMarksWithFilter = () => {
    emit('getMarksWithFilter')
}

let filterStore = useFilterStore();

const resetFilter = () => {
    filterStore.resetFilter()
}
const toggleBtn = (value) => {
    console.log(filterStore.filter.service_names)
    if (filterStore.filter.service_names.includes(value)) {
        filterStore.removeServiceName(value);
    } else {
        filterStore.addServiceName(value);
    }
}
</script>

<template>
    <div class="filter"
        :class="{active: props.active}"
    >
        <div class="filter__header">
            <div class="filter__header-text">
                <div class="filter__header-title">
                    Фильтры
                </div>
                <div class="filter__header-reset"
                    @click="resetFilter()"
                >
                    Сбросить
                </div>
            </div>
            <div class="filter__header-close"
                @click="closeFilter()"
            >
                <img src="@/assets/images/icons/close.svg">
            </div>
        </div>
        <div class="filter__list custom-scroll" v-if="props.activeObject === 'office'">
            <div class="filter__item _checkbox">
                <span class="title">
                    Доступно для маломобильных граждан
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.limited">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__item _checkbox">
                <span class="title">
                    Работает в выходные
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.weekend">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__item _money">
                <span class="title">
                    Снять
                </span>
                <div class="btns">
                    <button>
                        <img src="@/assets/images/icons/rub.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/dollar.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/evro.svg">
                    </button>
                </div>
                <p>
                    до
                </p>
                <input value="1 000 000">
            </div>
            <div class="filter__item _money">
                <span class="title">
                    Внести
                </span>
                <div class="btns">
                    <button>
                        <img src="@/assets/images/icons/rub.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/dollar.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/evro.svg">
                    </button>
                </div>
                <p>
                    до
                </p>
                <input value="1 000 000">
            </div>
            <div class="filter__item _money">
                <span class="title">
                    Перевести
                </span>
                <div class="btns">
                    <button>
                        <img src="@/assets/images/icons/rub.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/dollar.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/evro.svg">
                    </button>
                </div>
                <p>
                    до
                </p>
                <input value="1 000 000">
            </div>
            <div class="filter__item _select">
                <span class="title">
                    Карты
                </span>
                <div class="filter__select-list">
                    <button
                        :class="{btn_blue: filterStore.filter.service_names.includes('DebitCardIssuance')}"
                        @click="toggleBtn('DebitCardIssuance')"
                    >
                        Дебетовые
                    </button>
                    <button
                        :class="{btn_blue: filterStore.filter.service_names.includes('CreditCardIssuance')}"
                        @click="toggleBtn('CreditCardIssuance')"
                    >
                        Кредитные
                    </button>
                </div>
            </div>
            <div class="filter__item _select">
                <span class="title">
                    Инвестиции
                </span>
                <div class="filter__select-list">
                    <button
                        :class="{btn_blue: filterStore.filter.service_names.includes('BrokerService')}"
                        @click="toggleBtn('BrokerService')"
                    >
                        Брокерское обслуживание
                    </button>
                    <button
                        :class="{btn_blue: filterStore.filter.service_names.includes('TrustManagement')}"
                        @click="toggleBtn('TrustManagement')"
                    >
                        Доверительное управление
                    </button>
                </div>
            </div>
            <div class="filter__item _select">
                <span class="title">
                    Займы
                </span>
                <div class="filter__select-list">
                    <button
                        :class="{btn_blue: filterStore.filter.service_names.includes('CreditFL')}"
                        @click="toggleBtn('CreditFL')"
                    >
                        Кредит
                    </button>
                    <button
                        :class="{btn_blue: filterStore.filter.service_names.includes('CarCredit')}"
                        @click="toggleBtn('CarCredit')"
                    >
                        Автокредит
                    </button>
                    <button
                        :class="{btn_blue: filterStore.filter.service_names.includes('Mortgage')}"
                        @click="toggleBtn('Mortgage')"
                    >
                        Ипотека
                    </button>
                </div>
            </div>
            <div class="filter__item _select">
                <span class="title">
                    Услуги
                </span>
                <div class="filter__select-list">
                    <button
                        :class="{btn_blue: filterStore.filter.service_names.includes('ConversionFL')}"
                        @click="toggleBtn('ConversionFL')"
                    >
                        Конвертация валюты
                    </button>
                    <button
                        :class="{btn_blue: filterStore.filter.service_names.includes('PaymentForServicesFL')}"
                        @click="toggleBtn('PaymentForServicesFL')"
                    >
                        Оплата услуг
                    </button>
                </div>
            </div>
            <div class="filter__item _checkbox">
                <span class="title">
                    Аренда сейфовых ячеек
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.rent_boxes">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__item _checkbox">
                <span class="title">
                    Регистрация биометрии
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.biometric_registration">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__item _checkbox">
                <span class="title">
                    Sim карты ВТБ мобайл ИП
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.sims">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__item _checkbox">
                <span class="title">
                    Регистрация ИП
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.registration_ip">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__btns">
                <button class="filter__btn btn btn_blue"
                    @click="getMarksWithFilter()"
                >
                    Применить
                </button>
                <button class="filter__btn btn"
                    @click="closeFilter()"
                >
                    Отмена
                </button>
            </div>
        </div>
        <div class="filter__list custom-scroll" v-if="props.activeObject === 'atm'">
            <div class="filter__item _checkbox">
                <span class="title">
                    Доступно круглосуточно
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.all_day">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__item _checkbox">
                <span class="title">
                    Доступно для маломобильных
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.wheelchair">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__item _checkbox">
                <span class="title">
                    Оборудование для слабовидящих
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.blind">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__item _money">
                <span class="title">
                    Снять
                </span>
                <div class="btns">
                    <button>
                        <img src="@/assets/images/icons/rub.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/dollar.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/evro.svg">
                    </button>
                </div>
                <p>
                    до
                </p>
                <input value="1 000 000">
            </div>
            <div class="filter__item _money">
                <span class="title">
                    Внести
                </span>
                <div class="btns">
                    <button>
                        <img src="@/assets/images/icons/rub.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/dollar.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/evro.svg">
                    </button>
                </div>
                <p>
                    до
                </p>
                <input value="1 000 000">
            </div>
            <div class="filter__item _money">
                <span class="title">
                    Перевести
                </span>
                <div class="btns">
                    <button>
                        <img src="@/assets/images/icons/rub.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/dollar.svg">
                    </button>
                    <button>
                        <img src="@/assets/images/icons/evro.svg">
                    </button>
                </div>
                <p>
                    до
                </p>
                <input value="1 000 000">
            </div>
            <div class="filter__item _checkbox">
                <span class="title">
                    Снять наличные по QR-коду
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.qrRead">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__item _checkbox">
                <span class="title">
                    Платежи по QR-коду
                </span>
                <label class="switch">
                    <input type="checkbox" v-model="filterStore.filter.nfcForBankCards">
                    <span class="slider round"></span>
                </label>
            </div>
            <div class="filter__btns">
                <button class="filter__btn btn btn_blue"
                    @click="getMarksWithFilter()"
                >
                    Применить
                </button>
                <button class="filter__btn btn"
                    @click="closeFilter()"
                >
                    Отмена
                </button>
            </div>
        </div>
    </div>
</template>

<style lang="scss" >
.filter {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    z-index: 2;
    padding: 20px 0px;
    background: #1E1E1E;
    flex-direction: column;
    top: 100%;
    transition: .4s;
    display: flex;
    &.active {
        top: 0;
    }
    &__list {
        overflow: auto;
        flex: 1;
        padding: 0 20px;
    }
    &__header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 30px;
        width: 100%;
        margin-bottom: 10px;
        padding: 0 20px;
    }

    &__header-text {
        display: flex;
        align-items: flex-end;
        gap: 30px;
    }

    &__header-title {
        color: #FFF;
        font-family: VTB Group UI;
        font-size: 20px;
        font-style: normal;
        font-weight: 500;
        line-height: 120%; /* 24px */
    }

    &__header-reset {
        color: var(--absolute-input-gray, #ACB6C3);
        font-family: VTB Group UI;
        font-size: 12px;
        font-style: normal;
        font-weight: 400;
        line-height: 120%; /* 14.4px */
        cursor: pointer;
    }

    &__header-close {
        cursor: pointer;
    }

    &__item {
        &:not(:last-child) {
            margin-bottom: 30px;
        }
        &._checkbox {
            display: flex;
            justify-content: space-between;
            gap: 10px;
        }
        &._money {
            display: flex;
            align-items: center;
            color: #fff;
            justify-content: space-between;
            & .title {
                margin-right: 15px;
            }
            & .btns {
                display: flex;
                gap: 5px;
            }
            & button {
                display: flex;
                justify-content: center;
                align-items: center;
                width: 34px;
                height: 34px;
                border-radius: 8px;
                background: #2F3441;
            }
            & input {
                padding: 9px 10px;
                border-radius: 8px;
                background: var(--absolute-black, #2F3441);
                max-width: 90px;
                color: #FFF;
                font-family: VTB Group UI;
                font-size: 14px;
                font-style: normal;
                font-weight: 400;
                line-height: 140%; /* 19.6px */
            }
            & p {
                color: #FFF;
                font-size: 14px;
                font-style: normal;
                font-weight: 400;
                line-height: 140%; /* 19.6px */
            }
        }
        &._select {
            & button {
                border-radius: 30px;
                border: 1px solid #4789EB;
                color: var(--absolute-100, #FFF);
                font-family: VTB Group UI;
                font-size: 14px;
                font-style: normal;
                font-weight: 400;
                line-height: 140%; /* 19.6px */
                padding: 7px 10px;
            }
            & .title {
                margin-bottom: 20px;
                display: block;
            }
        }
        & .title {
            color: #FFF;
            font-family: VTB Group UI;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 120%; /* 19.2px */
        }
    }
    &__select-list {
        display: flex;
        gap: 15px;
        flex-wrap: wrap;
    }
    &__btns {
        display: flex;
        justify-content: center;
        gap: 10px;
        & .btn {
            flex: 0 0 126px;
        }
        & .btn_blue {
            flex: 1 ;
        }
    }
}





/* The switch - the box around the slider */
.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 26px;
  min-width: 60px;
}

/* Hide default HTML checkbox */
.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

/* The slider */
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #6B7683;
  -webkit-transition: .4s;
  transition: .4s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 0px;
  bottom: 0px;
  background-color: white;
  -webkit-transition: .4s;
  transition: .4s;
}

input:checked + .slider {
  background-color: #2196F3;
}

input:focus + .slider {
  box-shadow: 0 0 1px #2196F3;
}

input:checked + .slider:before {
  -webkit-transform: translateX(34px);
  -ms-transform: translateX(34px);
  transform: translateX(34px);
}

/* Rounded sliders */
.slider.round {
  border-radius: 34px;
}

.slider.round:before {
  border-radius: 50%;
}
</style>