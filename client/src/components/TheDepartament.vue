<script setup>
import { defineProps, defineEmits, computed } from 'vue';

const emit = defineEmits(['openInfoDepartament'])

const props = defineProps({
    departament: {
        type: Object,
    }
})

const openInfoDepartament = () => {
    emit('openInfoDepartament')
}
const kilometers = computed(() => {
    if (props.departament.distance) {
        const kilometers = props.departament.distance / 1000;
        return kilometers >= 10 ? kilometers.toFixed(1) + ' км.': kilometers + ' км.';
    } else {
        return '';
    }
});
</script>

<template>
    <div class="departament"
        @click="openInfoDepartament"
    >
        <div class="departament__icon">
            <img src="@/assets/images/icons/mark.svg">
        </div>
        <div class="departament__text">
            <div class="departament__address">
                г. Москва, Ленинский пр-т, д. 34/1
                {{ props.departament.address }}
            </div>
            <div class="departament__info">
                <span v-if="(props.departament.load <= 20 && props.departament.load >= 0) " style="color: green;">
                    не заполнено
                </span>
                <span v-if="(props.departament.load <= 40 && props.departament.load > 20)" style="color: orange;">
                    заполнено
                </span>
                <span v-if="(props.departament.load >= 40)" style="color: #CA181F;">
                    переполнено
                </span>
            </div>
        </div>
        <div class="departament__distance">
            {{ kilometers }}
        </div>
    </div>
</template>

<style lang="scss" scoped>
.departament {
    display: flex;
    align-items: center;
    padding: 12px;
    cursor: pointer;
    border: 1px solid transparent;
    border-radius: 12px;
    transition: .2s;
    &:hover {
        background: #96969675;
    }
    &__icon {
        margin-right: 16px;
    }

    &__text {
        margin-right: 6px;
    }

    &__address {
        color: var(--absolute-100, #FFF);
        font-family: VTB Group UI;
            font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 140%;
        margin-bottom: 10px;
    }

    &__info {
        color: #FFF;
        font-family: Inter;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        display: flex;
        align-items: center;
        gap: 5px;
        & .point {
            width: 5px;
            height: 5px;
            background: #fff;
            border-radius: 50%;
            margin-top: 3px;
        }
        & span {
            display: block;
        }
    }

    &__distance {
        color: #FFF;
        font-family: VTB Group UI;
            font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 140%; /* 25.2px */
        white-space: nowrap;
    }
}

</style>